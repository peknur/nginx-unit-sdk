package route

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutesMarshall(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/config.json")
	assert.NoError(t, err)

	var traverseMounts, followSymlinks bool
	actual, err := json.MarshalIndent(
		[]Config{
			{
				Match: &Match{
					URI:    "/admin/*",
					Scheme: "https",
					Arguments: map[string]string{
						"mode":   "strict",
						"access": "!raw",
					},
					Cookies: map[string]string{
						"user_role": "admin",
					},
				},
				Action: &Action{
					Pass: "applications/cms",
				},
			},
			{
				Match: &Match{
					Host: []string{
						"blog.example.com",
						"blog.*.org",
					},
					Source: []string{"*:8000-9000"},
				},
				Action: &Action{
					Pass: "applications/blogs/core",
				},
			},
			{
				Match: &Match{
					Host:   []string{"example.com"},
					Source: []string{"127.0.0.1-127.0.0.254:8080-8090"},
					URI:    "/chat/*",
					Query: []string{
						"en-CA",
						"en-IE",
						"en-IN",
						"en-UK",
						"en-US",
					},
				},
				Action: &Action{
					Pass: "applications/chat",
				},
			},
			{
				Match: &Match{
					Host: []string{"extwiki.example.com"},
				},
				Action: &Action{
					Pass: "applications/wiki/external",
				},
			},
			{
				Match: &Match{
					URI: "/legacy/*",
				},
				Action: &Action{
					Return:   301,
					Location: "https://legacy.example.com",
				},
			},

			{
				Match: &Match{
					Scheme: "http",
				},
				Action: &Action{
					Proxy: "http://127.0.0.1:8080",
				},
			},
			{
				Action: &Action{
					Share: []string{
						"/www/$host$uri",
						"/www/global_static$uri",
					},
					Chroot:         "/www/data/$host/",
					TraverseMounts: &traverseMounts,
					FollowSymlinks: &followSymlinks,
					Types: []string{
						"image/*",
						"video/*",
						"application/json",
					},
					Fallback: map[string]string{
						"proxy": "http://127.0.0.1:9000",
					},
				},
			},
		},
		"", "  ")
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}
