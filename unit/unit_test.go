package unit_test

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/peknur/nginx-unit-sdk/unit/client"
	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/application"
	"github.com/peknur/nginx-unit-sdk/unit/config/listener"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
	"github.com/peknur/nginx-unit-sdk/unit/config/settings"
	"github.com/peknur/nginx-unit-sdk/unit/config/upstream"
	"github.com/peknur/nginx-unit-sdk/unit/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const defaultUnitURL = "http://127.0.0.1:8080"

var svc service.Service

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
	client, err := client.NewClient(URL)
	if err != nil {
		log.Fatal(err)
	}
	svc = service.New(client)
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
		Applications: config.Applications{},
		AccessLog:    "/var/log/unit.log",
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
			ClientIP: &listener.ClientIP{
				Header: "X-Demo",
				Source: []string{"127.0.0.1"},
			},
		},
	}))

	assert.NoError(t, svc.CreateListener(ctx, "*:80", listener.Config{
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

func Example() {
	client, err := client.New("http://127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	svc = service.New(client)
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
		AccessLog: "/var/log/unit.log",
	}

	if err := svc.CreateConfig(context.Background(), cfg); err != nil {
		log.Fatal(err)
	}
}
