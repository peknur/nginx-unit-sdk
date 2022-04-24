package service

import (
	"context"
	"path"

	"github.com/peknur/nginx-unit-sdk/unit/config"
	"github.com/peknur/nginx-unit-sdk/unit/config/application"
)

func (s *Service) Applications(ctx context.Context) (config.Applications, error) {
	c := config.Applications{}
	return c, s.client.Get(ctx, applicationsPath, &c)
}

func (s *Service) CreateApplications(ctx context.Context, c config.Applications) error {
	return s.client.Put(ctx, applicationsPath, c)
}

func (s *Service) CreateApplication(ctx context.Context, name string, c application.Config) error {
	return s.client.Put(ctx, path.Join(applicationsPath, name), c)
}

func (s *Service) DeleteApplication(ctx context.Context, name string) error {
	return s.client.Delete(ctx, path.Join(applicationsPath, name))
}
