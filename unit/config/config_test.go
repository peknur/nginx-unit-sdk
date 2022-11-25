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
		AccessLog: AccessLog{
			Path:   "/var/log/access.log",
			Format: "$remote_addr - - [$time_local] \"$request_line\" $status $body_bytes_sent \"$header_referer\" \"$header_user_agent\"",
		},
	})
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}
