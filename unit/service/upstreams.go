package service

import (
	"context"
	"path"

	"github.com/peknur/nginx-unit-sdk/unit"
	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/upstream"
)

func (s *Service) Upstreams(ctx context.Context) (config.Upstreams, error) {
	c := config.Upstreams{}
	return c, s.client.Get(ctx, unit.UpstreamsPath, &c)
}

func (s *Service) CreateUpstreams(ctx context.Context, c config.Upstreams) error {
	return s.client.Put(ctx, unit.UpstreamsPath, c)
}

func (s *Service) CreateUpstream(ctx context.Context, name string, c upstream.Config) error {
	return s.client.Put(ctx, path.Join(unit.UpstreamsPath, name), c)
}

func (s *Service) DeleteUpstream(ctx context.Context, name string) error {
	return s.client.Delete(ctx, path.Join(unit.UpstreamsPath, name))
}
