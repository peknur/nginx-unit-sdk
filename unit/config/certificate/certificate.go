package certificate

import (
	"fmt"
	"strings"
	"time"
)

// Sample Sep 18 19:46:19 2018 GMT
const CertificateTimeLayout = "Jan 02 15:04:05 2006 GMT"

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format(CertificateTimeLayout))), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	ts, err := time.Parse(CertificateTimeLayout, s)
	*t = Time(ts)
	return err
}

type Config struct {
	Key   string  `json:"key,omitempty"`
	Chain []Chain `json:"chain,omitempty"`
}

type Chain struct {
	Subject  Subject  `json:"subject,omitempty"`
	Issuer   Issuer   `json:"issuer,omitempty"`
	Validity Validity `json:"validity,omitempty"`
}

type Subject struct {
	CommonName      string   `json:"common_name,omitempty"`
	AltNames        []string `json:"alt_names,omitempty"`
	Country         string   `json:"country,omitempty"`
	StateOrProvince string   `json:"state_or_province,omitempty"`
	Organization    string   `json:"organization,omitempty"`
}

type Issuer struct {
	CommonName      string `json:"common_name,omitempty"`
	Country         string `json:"country,omitempty"`
	StateOrProvince string `json:"state_or_province,omitempty"`
	Organization    string `json:"organization,omitempty"`
}

type Validity struct {
	Since Time `json:"since,omitempty"`
	Until Time `json:"until,omitempty"`
}
