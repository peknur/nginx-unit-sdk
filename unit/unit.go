package unit

import (
	"context"

	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/application"
	"github.com/peknur/nginx-unit-sdk/unit/config/listener"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
	"github.com/peknur/nginx-unit-sdk/unit/config/upstream"
)

const (
	CertificatesPath string = "certificates"
	ConfigPath       string = "config"
	ApplicationsPath string = "config/applications"
	ListenersPath    string = "config/listeners"
	RoutesPath       string = "config/routes"
	UpstreamsPath    string = "config/upstreams"
	SettingsPath     string = "config/settings"
)

type Client interface {
	// Get returns the entity at the request URI.
	Get(ctx context.Context, path string, v interface{}) error
	// Put replaces the entity at the request URI.
	Put(ctx context.Context, path string, v interface{}) error
	// PutBinary replaces the entity at the request URI with data.
	PutBinary(ctx context.Context, path string, data []byte) error
	// Post updates the array at the request URI.
	Post(ctx context.Context, path string, v interface{}) error
	// Delete deletes the entity at the request URI.
	Delete(ctx context.Context, path string) error
}

type Service interface {
	// Applications
	Applications(ctx context.Context) (config.Applications, error)
	CreateApplications(ctx context.Context, c config.Applications) error
	CreateApplication(ctx context.Context, name string, c application.Config) error
	DeleteApplication(ctx context.Context, name string) error

	// Config
	Config(ctx context.Context) (config.Config, error)
	CreateConfig(ctx context.Context, c config.Config) error
	DeleteConfig(ctx context.Context) error

	// Listeners
	Listeners(ctx context.Context) (config.Listeners, error)
	CreateListener(ctx context.Context, name string, c listener.Config) error
	CreateListeners(ctx context.Context, c config.Listeners) error
	DeleteListener(ctx context.Context, name string) error

	// Routes
	Routes(ctx context.Context) (config.Routes, error)
	CreateRoute(ctx context.Context, name string, c []route.Config) error
	CreateRoutes(ctx context.Context, c config.Routes) error
	AppendRoute(ctx context.Context, name string, c route.Config) error
	DeleteRoute(ctx context.Context, name string) error

	// Upstreams
	Upstreams(ctx context.Context) (config.Upstreams, error)
	CreateUpstream(ctx context.Context, name string, c upstream.Config) error
	CreateUpstreams(ctx context.Context, c config.Upstreams) error
	DeleteUpstream(ctx context.Context, name string) error

	// Certificates
	Certificates(ctx context.Context) (config.Certificates, error)
	CreateCertificate(ctx context.Context, name string, bundle []byte) error
	DeleteCertificate(ctx context.Context, name string) error

	// Settings
	Settings(ctx context.Context) (config.Settings, error)
	CreateSettings(ctx context.Context, c config.Settings) error
	DeleteSettings(ctx context.Context) error
}
