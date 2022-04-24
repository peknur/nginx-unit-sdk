package service

import (
	"context"
)

const (
	certificatesPath string = "certificates"
	configPath       string = "config"
	applicationsPath string = "config/applications"
	listenersPath    string = "config/listeners"
	routesPath       string = "config/routes"
	upstreamsPath    string = "config/upstreams"
	settingsPath     string = "config/settings"
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

type Service struct {
	client Client
}

func New(client Client) *Service {
	return &Service{client: client}
}
