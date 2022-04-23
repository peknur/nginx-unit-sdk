package settings

type HTTP struct {
	HeaderReadTimeout   int   `json:"header_read_timeout,omitempty"`
	BodyReadTimeout     int   `json:"body_read_timeout,omitempty"`
	SendTimeout         int   `json:"send_timeout,omitempty"`
	IdleTimeout         int   `json:"idle_timeout,omitempty"`
	MaxBodySize         int   `json:"max_body_size,omitempty"`
	Static              *Mime `json:"static,omitempty"`
	DiscardUnsafeFields *bool `json:"discard_unsafe_fields,omitempty"`
}

type Mime struct {
	MimeTypes map[string][]string `json:"mime_types,omitempty"`
}
