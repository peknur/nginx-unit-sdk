package upstream

type Config struct {
	Servers map[string]Server `json:"servers,omitempty"`
}

type Server struct {
	Weight int `json:"weight,omitempty"`
}
