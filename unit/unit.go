package unit

import (
	"context"

	"github.com/peknur/nginx-unit-sdk/unit/certificate"
	"github.com/peknur/nginx-unit-sdk/unit/client"
	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/application"
	"github.com/peknur/nginx-unit-sdk/unit/config/listener"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
	"github.com/peknur/nginx-unit-sdk/unit/config/upstream"
	"github.com/peknur/nginx-unit-sdk/unit/service"
	"github.com/peknur/nginx-unit-sdk/unit/status"
)

// Service interface defines methods that can be used to interact with Unit instance.
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
	Certificates(ctx context.Context) (certificate.Certificates, error)
	CreateCertificate(ctx context.Context, name string, bundle []byte) error
	DeleteCertificate(ctx context.Context, name string) error

	// Settings
	Settings(ctx context.Context) (config.Settings, error)
	CreateSettings(ctx context.Context, c config.Settings) error
	DeleteSettings(ctx context.Context) error

	// Status
	Status(ctx context.Context) (status.Status, error)
}

// NewServiceFromURL creates new service instance using URL as client base URL.
func NewServiceFromURL(URL string) (Service, error) {
	c, err := client.NewClient(URL)
	if err != nil {
		return nil, err
	}
	return NewService(c), nil
}

// NewService creates new service instance.
func NewService(client service.Client) Service {
	return service.New(client)
}

type Unit struct {
	Certificates certificate.Certificates `json:"certificates,omitempty"`
	Config       config.Config            `json:"config,omitempty"`
}

type Certificates map[string]certificate.Config
