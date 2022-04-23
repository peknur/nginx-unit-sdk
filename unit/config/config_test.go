package config

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigMarshall(t *testing.T) {
	expected, err := ioutil.ReadFile("testdata/config.json")
	assert.NoError(t, err)

	actual, err := json.Marshal(Config{
		Settings:     Settings{},
		Listeners:    Listeners{},
		Routes:       Routes{},
		Applications: Applications{},
		Upstreams:    Upstreams{},
		AccessLog:    "/var/log/access.log",
	})
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}

func TestUnitMarshall(t *testing.T) {
	expected, err := ioutil.ReadFile("testdata/unit.json")
	assert.NoError(t, err)

	actual, err := json.Marshal(Unit{
		Certificates: Certificates{},
		Config: Config{
			Settings:     Settings{},
			Listeners:    Listeners{},
			Routes:       Routes{},
			Applications: Applications{},
			Upstreams:    Upstreams{},
			AccessLog:    "/var/log/access.log",
		},
	})
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}
