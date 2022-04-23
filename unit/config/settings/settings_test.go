package settings

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettingsMarshall(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/config.json")
	assert.NoError(t, err)

	var discardUnsafeFields bool
	actual, err := json.MarshalIndent(
		map[string]HTTP{
			"http": {
				HeaderReadTimeout: 10,
				BodyReadTimeout:   10,
				SendTimeout:       10,
				IdleTimeout:       120,
				MaxBodySize:       6291456,
				Static: &Mime{
					MimeTypes: map[string][]string{
						"text/plain": {
							".log",
							"README",
							"CHANGES",
						}},
				},
				DiscardUnsafeFields: &discardUnsafeFields,
			},
		},
		"", "  ")
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}
