package status

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusMarshall(t *testing.T) {
	expected, err := ioutil.ReadFile("testdata/status.json")
	assert.NoError(t, err)

	actual, err := json.Marshal(Status{
		Connections: Connections{Accepted: 1067, Active: 13, Idle: 4, Closed: 1050},
		Requests:    Requests{Total: 1307},
		Applications: map[string]Application{
			"wp": {
				Processes: ApplicationProcesses{
					Running:  14,
					Starting: 0,
					Idle:     4,
				},
				Requests: ApplicationRequests{
					Active: 10,
				},
			},
		},
	})
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}
