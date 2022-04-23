package service

import (
	"context"
	"path"

	"github.com/peknur/nginx-unit-sdk/unit"
	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/listener"
)

func (s *Service) Listeners(ctx context.Context) (config.Listeners, error) {
	c := config.Listeners{}
	return c, s.client.Get(ctx, unit.ListenersPath, &c)
}

func (s *Service) CreateListeners(ctx context.Context, c config.Listeners) error {
	return s.client.Put(ctx, unit.ListenersPath, c)
}

func (s *Service) CreateListener(ctx context.Context, name string, c listener.Config) error {
	return s.client.Put(ctx, path.Join(unit.ListenersPath, name), c)
}

func (s *Service) DeleteListener(ctx context.Context, name string) error {
	return s.client.Delete(ctx, path.Join(unit.ListenersPath, name))
}
