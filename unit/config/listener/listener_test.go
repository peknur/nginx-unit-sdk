package listener

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListenersMarshall(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/config.json")
	assert.NoError(t, err)

	actual, err := json.MarshalIndent(
		map[string]Config{
			"*:8000": {
				Pass: "routes",
				TLS: &TLS{
					Certificate: []string{
						"example.com",
						"example.com",
					},
					ConfCommands: map[string]string{
						"ciphersuites": "TLS_CHACHA20_POLY1305_SHA256",
					},
					Session: &Session{
						CacheSize: 10240,
						Timeout:   60,
						Tickets: []string{
							"k5qMHi7IMC7ktrPY3lZ+sL0Zm8oC0yz6re+y/zCj0H0/sGZ7yPBwGcb77i5vw6vCx8vsQDyuvmFb6PZbf03Auj/cs5IHDTYkKIcfbwz6zSU=",
							"3Cy+xMFsCjAek3TvXQNmCyfXCnFNAcAOyH5xtEaxvrvyyCS8PJnjOiq2t4Rtf/Gq",
							"8dUI0x3LRnxfN0miaYla46LFslJJiBDNdFiPJdqr37mYQVIzOWr+ROhyb1hpmg/QCM2qkIEWJfrJX3I+rwm0t0p4EGdEVOXQj7Z8vHFcbiA=",
						},
					},
				},
			},
			"127.0.0.1:8001": {
				Pass: "applications/drive",
			},
			"*:8080": {
				Pass: "upstreams/rr-lb",
				ClientIP: &ClientIP{
					Header: "X-Forwarded-For",
					Source: []string{
						"192.168.0.0.0/16",
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
