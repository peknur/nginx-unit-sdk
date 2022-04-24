package service

import (
	"context"

	"github.com/peknur/nginx-unit-sdk/unit/config"
)

func (s *Service) Settings(ctx context.Context) (config.Settings, error) {
	c := config.Settings{}
	return c, s.client.Get(ctx, settingsPath, &c)
}

func (s *Service) CreateSettings(ctx context.Context, c config.Settings) error {
	return s.client.Put(ctx, settingsPath, c)
}

func (s *Service) DeleteSettings(ctx context.Context) error {
	return s.client.Delete(ctx, settingsPath)
}
