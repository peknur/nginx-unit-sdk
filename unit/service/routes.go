package service

import (
	"context"
	"path"

	"github.com/peknur/nginx-unit-sdk/unit"
	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/route"
)

func (s *Service) Routes(ctx context.Context) (config.Routes, error) {
	c := config.Routes{}
	return c, s.client.Get(ctx, unit.RoutesPath, &c)
}

func (s *Service) CreateRoutes(ctx context.Context, c config.Routes) error {
	return s.client.Put(ctx, unit.RoutesPath, c)
}

func (s *Service) CreateRoute(ctx context.Context, name string, c []route.Config) error {
	return s.client.Put(ctx, path.Join(unit.RoutesPath, name), c)
}

func (s *Service) AppendRoute(ctx context.Context, name string, c route.Config) error {
	return s.client.Post(ctx, path.Join(unit.RoutesPath, name), c)
}

func (s *Service) DeleteRoute(ctx context.Context, name string) error {
	return s.client.Delete(ctx, path.Join(unit.RoutesPath, name))
}
