package service

import (
	"context"

	"github.com/peknur/nginx-unit-sdk/unit/config"
)

func (s *Service) Config(ctx context.Context) (config.Config, error) {
	c := config.Config{}
	return c, s.client.Get(ctx, configPath, &c)
}

func (s *Service) CreateConfig(ctx context.Context, c config.Config) error {
	return s.client.Put(ctx, configPath, c)
}

func (s *Service) DeleteConfig(ctx context.Context) error {
	return s.client.Delete(ctx, configPath)
}
