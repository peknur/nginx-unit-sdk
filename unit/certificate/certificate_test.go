package certificate

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCertificateMarshall(t *testing.T) {
	t.Parallel()
	expected, err := ioutil.ReadFile("testdata/config.json")
	assert.NoError(t, err)
	certSince, err := time.Parse(CertificateTimeLayout, "Sep 18 19:46:19 2018 GMT")
	if err != nil {
		t.Fatal(err)
	}
	certUntil, err := time.Parse(CertificateTimeLayout, "Jun 15 19:46:19 2021 GMT")
	if err != nil {
		t.Fatal(err)
	}
	actual, err := json.MarshalIndent(
		map[string]Config{
			"example.com": {
				Key: "RSA (4096 bits)",
				Chain: []Chain{
					{
						Subject: Subject{
							CommonName: "example.com",
							AltNames: []string{
								"example.com",
								"www.example.com",
							},
							Country:         "US",
							StateOrProvince: "CA",
							Organization:    "Acme, Inc.",
						},
						Issuer: Issuer{
							CommonName:      "intermediate.ca.example.com",
							Country:         "US",
							StateOrProvince: "CA",
							Organization:    "Acme Certification Authority",
						},
						Validity: Validity{
							Since: Time(certSince),
							Until: Time(certUntil),
						},
					},
					{
						Subject: Subject{
							CommonName:      "intermediate.ca.example.com",
							Country:         "US",
							StateOrProvince: "CA",
							Organization:    "Acme Certification Authority",
						},
						Issuer: Issuer{
							CommonName:      "root.ca.example.com",
							Country:         "US",
							StateOrProvince: "CA",
							Organization:    "Acme Root Certification Authority",
						},
						Validity: Validity{
							Since: Time(certSince),
							Until: Time(certUntil),
						},
					},
				},
			},
		},
		" ", " ")
	assert.NoError(t, err)
	if !assert.JSONEq(t, string(expected), string(actual)) {
		t.Log(string(actual))
	}
}
