package listener

type Config struct {
	Pass     string    `json:"pass,omitempty"`
	ClientIP *ClientIP `json:"client_ip,omitempty"`
	TLS      *TLS      `json:"tls,omitempty"`
}

type ClientIP struct {
	Header string   `json:"header,omitempty"`
	Source []string `json:"source,omitempty"`
}

type TLS struct {
	Certificate  []string          `json:"certificate,omitempty"`
	ConfCommands map[string]string `json:"conf_commands,omitempty"`
	Session      *Session          `json:"session,omitempty"`
}

type Session struct {
	CacheSize int      `json:"cache_size,omitempty"`
	Timeout   int      `json:"timeout,omitempty"`
	Tickets   []string `json:"tickets,omitempty"`
}
