package listener

type Config struct {
	Pass      string     `json:"pass,omitempty"`
	Forwarded *Forwarded `json:"forwarded,omitempty"`
	TLS       *TLS       `json:"tls,omitempty"`
}

type Forwarded struct {
	ClientIP  string   `json:"client_ip,omitempty"`
	Recursive bool     `json:"recursive,omitempty"`
	Protocol  string   `json:"protocol,omitempty"`
	Header    string   `json:"header,omitempty"`
	Source    []string `json:"source,omitempty"`
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
