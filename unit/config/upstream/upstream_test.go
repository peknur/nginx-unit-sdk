package upstream

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpstreamsMarshall(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/config.json")
	assert.NoError(t, err)

	actual, err := json.MarshalIndent(
		map[string]Config{
			"rr-lb": {
				Servers: map[string]Server{
					"192.168.1.100:8080": {},
					"192.168.1.101:8080": {
						Weight: 2,
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
