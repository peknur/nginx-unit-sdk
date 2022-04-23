package application

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationMarshall(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/config.json")
	assert.NoError(t, err)
	actual, err := json.MarshalIndent(
		map[string]Config{
			"chat": {
				Type:             TypeExternal,
				Executable:       "bin/chat_app",
				Group:            "www-chat",
				User:             "www-chat",
				WorkingDirectory: "/www/chat/",
				Isolation: &Isolation{
					Namespaces: &Namespaces{
						Cgroup:     false,
						Credential: true,
						Mount:      false,
						Network:    false,
						PID:        false,
						Uname:      false,
					},
					Uidmap: &[]UIDMap{
						{
							Host:      1000,
							Container: 0,
							Size:      1000,
						},
					},
					Gidmap: &[]GIDMap{
						{
							Host:      1000,
							Container: 0,
							Size:      1000,
						},
					},
					Automount: &Automount{
						LanguageDeps: false,
						Procfs:       false,
						Tmpfs:        false,
					},
				},
			},
			"cms": {
				Type:             TypeRuby,
				Script:           "/www/cms/main.ru",
				WorkingDirectory: "/www/cms/",
				Hooks:            "hooks.rb",
			},
			"drive": {
				Type:             TypePerl,
				Script:           "app.psgi",
				Threads:          2,
				ThreadStackSize:  4096,
				WorkingDirectory: "/www/drive/",
				Processes: &Processes{
					Max:         10,
					Spare:       5,
					IdleTimeout: 20,
				},
			},
			"wiki": {
				Type:     TypePython,
				Protocol: "asgi",
				Targets: map[string]Target{
					"internal": {
						Module: "internal.asgi",
					},
					"external": {
						Module: "external.asgi",
					},
				},
				Environment: map[string]string{
					"DJANGO_SETTINGS_MODULE": "wiki.settings.prod",
					"DB_ENGINE":              "django.db.backends.postgresql",
					"DB_NAME":                "wiki",
					"DB_HOST":                "127.0.0.1",
					"DB_PORT":                "5432",
				},
				Path: []string{"/www/wiki/"},
				Processes: &Processes{
					Max: 10,
				},
			},
		},
		" ", " ")
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}

func TestJava(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/java.json")
	assert.NoError(t, err)
	cfg := Config{
		Type:   TypeJava,
		Webapp: "/www/store/store.war",
		Classpath: []string{
			"/www/store/lib/store-2.0.0.jar",
		},
		Options: []string{
			"-Dlog_path=/var/log/store.log",
		},
	}
	actual, err := json.MarshalIndent(cfg, " ", " ")
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
	c := Config{}
	err = json.Unmarshal(expected, &c)
	assert.NoError(t, err)
	assert.Equal(t, cfg, c)
}

func TestPHP(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/php.json")
	assert.NoError(t, err)
	cfg := Config{
		Type: TypePHP,
		Targets: map[string]Target{
			"admin": {
				Root:   "/www/blogs/admin/",
				Script: "index.php",
			},
			"core": {
				Root: "/www/blogs/scripts/",
			},
		},
		Limits: &Limits{
			Timeout:  10,
			Requests: 1000,
		},
		Options: PHPOptions{
			File: "/etc/php.ini",
			Admin: map[string]string{
				"memory_limit":    "256M",
				"variables_order": "EGPCS",
				"expose_php":      "0",
			},
			User: map[string]string{
				"display_errors": "0",
			},
		},
		Processes: &Processes{
			Max: 4,
		},
	}
	actual, err := json.MarshalIndent(cfg, " ", " ")
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
	c := Config{}
	err = json.Unmarshal(expected, &c)
	assert.NoError(t, err)
	assert.Equal(t, cfg, c)
}
