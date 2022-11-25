package unit_test

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"

	"github.com/peknur/nginx-unit-sdk/unit"
	"github.com/peknur/nginx-unit-sdk/unit/client"
	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/application"
	"github.com/peknur/nginx-unit-sdk/unit/config/listener"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
	"github.com/peknur/nginx-unit-sdk/unit/config/settings"
	"github.com/peknur/nginx-unit-sdk/unit/config/upstream"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultUnitURL string = "http://127.0.0.1:8080"
	appURL         string = "http://127.0.0.1:8081"
)

var svc unit.Service

func TestMain(m *testing.M) {
	flag.Parse()
	if testing.Short() {
		log.Println("skipping integration test")
		return
	}
	URL := defaultUnitURL
	if u := os.Getenv("TEST_UNIT_URL"); u != "" {
		URL = u
	}
	var err error
	svc, err = unit.NewServiceFromURL(URL)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestConfig(t *testing.T) {
	ctx := context.Background()
	cfg := config.Config{
		Settings: config.Settings{
			HTTP: settings.HTTP{
				HeaderReadTimeout: 10,
				BodyReadTimeout:   10,
			},
		},
		Listeners: config.Listeners{
			"*:8081": listener.Config{
				Pass: "routes/main",
			},
		},
		Routes: config.Routes{
			"main": []route.Config{
				{
					Match: &route.Match{
						Host: []string{"example.com"},
					},
					Action: &route.Action{
						Return: http.StatusNotFound,
					},
				},
			},
		},
		Applications: config.Applications{},
		AccessLog: config.AccessLog{
			Path:   "/var/log/access.log",
			Format: "$remote_addr - - [$time_local] \"$request_line\" $status $body_bytes_sent \"$header_referer\" \"$header_user_agent\"",
		},
	}

	assert.NoError(t, svc.CreateConfig(ctx, cfg))
	c, err := svc.Config(ctx)
	assert.NoError(t, err)
	assert.Equal(t, cfg, c)
	testListeners(ctx, t)
	testRoutes(ctx, t)
	testUpstream(ctx, t)
	testSettings(ctx, t)
	testApplications(ctx, t)
	testCertificates(ctx, t)
	testCertificates(ctx, t)
	assert.NoError(t, svc.DeleteConfig(ctx))
}

func testListeners(ctx context.Context, t *testing.T) {
	assert.NoError(t, svc.CreateListeners(ctx, config.Listeners{
		"*:443": listener.Config{
			Pass: "routes/main",
			Forwarded: &listener.Forwarded{
				Recursive: true,
				Protocol:  "X-Forwarded-Proto",
				Source:    []string{"127.0.0.1"},
			},
		},
	}))

	assert.NoError(t, svc.CreateListener(ctx, "*:8081", listener.Config{
		Pass: "routes/main",
	}))

	assert.NoError(t, svc.DeleteListener(ctx, "*:443"))

	assert.NoError(t, svc.CreateListener(ctx, "*:443", listener.Config{
		Pass: "routes/main",
	}))

	listeners, err := svc.Listeners(ctx)
	assert.NoError(t, err)
	assert.Len(t, listeners, 2)
}

func testRoutes(ctx context.Context, t *testing.T) {
	assert.NoError(t, svc.CreateRoute(ctx, "sec", []route.Config{
		{
			Match: &route.Match{
				Host: []string{"example.net"},
			},
			Action: &route.Action{
				Return: http.StatusNotFound,
			},
		},
	}))

	assert.NoError(t, svc.AppendRoute(context.Background(), "sec", route.Config{
		Match: &route.Match{
			Host: []string{"example.org"},
		},
		Action: &route.Action{
			Return: http.StatusNotFound,
		},
	}))

	assert.NoError(t, svc.DeleteRoute(ctx, "sec"))

	assert.NoError(t, svc.CreateRoutes(context.Background(),
		config.Routes{
			"main": {{
				Match: &route.Match{
					Host: []string{"example.org"},
				},
				Action: &route.Action{
					Return: http.StatusNotFound,
				}},
			},
			"test2": {{
				Match: &route.Match{
					Host: []string{"example.org"},
				},
				Action: &route.Action{
					Return: http.StatusNotFound,
				}},
			},
		}))

	routes, err := svc.Routes(ctx)
	assert.NoError(t, err)
	assert.Len(t, routes, 2)
}

func testUpstream(ctx context.Context, t *testing.T) {
	assert.NoError(t, svc.CreateUpstreams(ctx, config.Upstreams{
		"test1": upstream.Config{
			Servers: map[string]upstream.Server{
				"127.0.0.1:8080": {
					Weight: 0,
				},
			},
		},
		"test2": upstream.Config{
			Servers: map[string]upstream.Server{
				"127.0.0.1:8080": {
					Weight: 0,
				},
			},
		},
	}))
	assert.NoError(t, svc.CreateUpstream(ctx, "test3", upstream.Config{
		Servers: map[string]upstream.Server{
			"127.0.0.1:8080": {
				Weight: 0,
			},
		},
	}))
	assert.NoError(t, svc.CreateUpstream(ctx, "test1", upstream.Config{
		Servers: map[string]upstream.Server{
			"127.0.0.1:8080": {
				Weight: 0,
			},
		},
	}))
	ups, err := svc.Upstreams(ctx)
	assert.NoError(t, err)
	assert.Len(t, ups, 3)
	assert.NoError(t, svc.DeleteUpstream(ctx, "test1"))

}

func testApplications(ctx context.Context, t *testing.T) {
	assert.NoError(t, svc.CreateApplications(ctx, config.Applications{
		"app1": {
			Type:             application.TypeGo,
			Executable:       "app",
			WorkingDirectory: "/apps/go",
			User:             "www-data",
			Group:            "www-data",
		},
		"app2": {
			Type:             application.TypeGo,
			Executable:       "app",
			WorkingDirectory: "/apps/go",
			User:             "www-data",
			Group:            "www-data",
		},
	},
	))
	assert.NoError(t, svc.CreateApplication(ctx, "go", application.Config{
		Type:       application.TypeGo,
		Executable: "app",
		Limits: &application.Limits{
			Timeout:  10,
			Requests: 100,
		},
		Processes: &application.Processes{
			Max:         1000,
			Spare:       10,
			IdleTimeout: 10,
		},
		WorkingDirectory: "/apps/go",
		User:             "www-data",
		Group:            "www-data",
	}))
	apps, err := svc.Applications(ctx)
	assert.NoError(t, err)
	assert.Len(t, apps, 3)
	assert.NoError(t, svc.CreateRoute(ctx, "main", []route.Config{
		{
			Match: &route.Match{
				Scheme: "http",
			},
			Action: &route.Action{
				Pass: "applications/app1",
			},
		},
	}))
	testStatus(ctx, t)
	assert.NoError(t, svc.DeleteApplication(ctx, "go"))
}

func testCertificates(ctx context.Context, t *testing.T) {
	bundle, err := ioutil.ReadFile("testdata/sample_certificate.pem")
	require.NoError(t, err)
	assert.NoError(t, svc.CreateCertificate(ctx, "c1", bundle))
	certs, err := svc.Certificates(ctx)
	assert.NoError(t, err)
	assert.Len(t, certs, 1)
	assert.NoError(t, svc.DeleteCertificate(ctx, "c1"))
}

func testSettings(ctx context.Context, t *testing.T) {
	assert.NoError(t, svc.DeleteSettings(ctx))
	assert.NoError(t, svc.CreateSettings(ctx, config.Settings{
		HTTP: settings.HTTP{
			HeaderReadTimeout: 20,
		},
	}))

	s, err := svc.Settings(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 20, s.HTTP.HeaderReadTimeout)
}

func testStatus(ctx context.Context, t *testing.T) {
	var wg sync.WaitGroup
	c := 50
	for i := 0; i < c; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := http.Get(appURL)
			assert.NoError(t, err)
		}()
	}
	wg.Wait()
	s, err := svc.Status(ctx)
	assert.NoError(t, err)
	assert.LessOrEqual(t, c, s.Requests.Total)
	assert.LessOrEqual(t, c, s.Connections.Accepted)
}

func Example() {
	client, err := client.New("http://127.0.0.1:8080", http.DefaultClient)
	if err != nil {
		log.Fatal(err)
	}
	svc := unit.NewService(client)
	cfg := config.Config{
		Settings: config.Settings{
			HTTP: settings.HTTP{
				HeaderReadTimeout: 10,
				BodyReadTimeout:   10,
			},
		},
		Listeners: config.Listeners{
			"*:80": listener.Config{
				Pass: "routes/main",
			},
		},
		Routes: config.Routes{
			"main": []route.Config{
				{
					Match: &route.Match{
						Host: []string{"example.com"},
					},
					Action: &route.Action{
						Return: http.StatusNotFound,
					},
				}},
		},
		AccessLog: config.AccessLog{
			Path:   "/var/log/access.log",
			Format: "$remote_addr - - [$time_local] \"$request_line\" $status $body_bytes_sent \"$header_referer\" \"$header_user_agent\"",
		},
	}

	if err := svc.CreateConfig(context.Background(), cfg); err != nil {
		log.Fatal(err)
	}
}

func ExampleNewServiceFromURL() {
	svc, err := unit.NewServiceFromURL("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	svc.Config(context.TODO())
}
