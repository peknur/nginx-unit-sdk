package config

import (
	"github.com/peknur/nginx-unit-sdk/unit/config/application"
	"github.com/peknur/nginx-unit-sdk/unit/config/certificate"
	"github.com/peknur/nginx-unit-sdk/unit/config/listener"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
	"github.com/peknur/nginx-unit-sdk/unit/config/settings"
	"github.com/peknur/nginx-unit-sdk/unit/config/upstream"
)

type Entity interface {
	Path() string
}

type Unit struct {
	Certificates Certificates `json:"certificates,omitempty"`
	Config       Config       `json:"config,omitempty"`
}

type Certificates map[string]certificate.Config

func (c *Certificates) Path() string {
	return "certificates"
}

type Config struct {
	Settings     Settings     `json:"settings"`
	Listeners    Listeners    `json:"listeners"`
	Routes       Routes       `json:"routes"`
	Applications Applications `json:"applications"`
	Upstreams    Upstreams    `json:"upstreams,omitempty"`
	AccessLog    string       `json:"access_log,omitempty"`
}

func (c *Config) Path() string {
	return "config"
}

type Applications map[string]application.Config

func (a *Applications) Path() string {
	return "config/applications"
}

type Listeners map[string]listener.Config

func (l *Listeners) Path() string {
	return "config/listeners"
}

type Routes map[string][]route.Config

func (r *Routes) Path() string {
	return "config/routes"
}

type Upstreams map[string]upstream.Config

func (u *Upstreams) Path() string {
	return "config/upstreams"
}

type Settings struct {
	HTTP settings.HTTP `json:"http,omitempty"`
}

func (s *Settings) Path() string {
	return "config/settings"
}
