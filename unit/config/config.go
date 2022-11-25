package config

import (
	"github.com/peknur/nginx-unit-sdk/unit/config/application"
	"github.com/peknur/nginx-unit-sdk/unit/config/listener"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
	"github.com/peknur/nginx-unit-sdk/unit/config/settings"
	"github.com/peknur/nginx-unit-sdk/unit/config/upstream"
)

type Config struct {
	Settings     Settings     `json:"settings"`
	Listeners    Listeners    `json:"listeners"`
	Routes       Routes       `json:"routes"`
	Applications Applications `json:"applications"`
	Upstreams    Upstreams    `json:"upstreams,omitempty"`
	AccessLog    AccessLog    `json:"access_log,omitempty"`
}

type AccessLog struct {
	Path   string `json:"path,omitempty"`
	Format string `json:"format,omitempty"`
}

type Applications map[string]application.Config

type Listeners map[string]listener.Config

type Routes map[string][]route.Config

type Upstreams map[string]upstream.Config

type Settings struct {
	HTTP settings.HTTP `json:"http,omitempty"`
}
