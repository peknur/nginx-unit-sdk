package service

import (
	"context"
	"path"

	"github.com/peknur/nginx-unit-sdk/unit/config"
)

func (s *Service) Certificates(ctx context.Context) (config.Certificates, error) {
	c := config.Certificates{}
	return c, s.client.Get(ctx, certificatesPath, &c)
}

func (s *Service) CreateCertificate(ctx context.Context, name string, bundle []byte) error {
	return s.client.PutBinary(ctx, path.Join(certificatesPath, name), bundle)
}

func (s *Service) DeleteCertificate(ctx context.Context, name string) error {
	return s.client.Delete(ctx, path.Join(certificatesPath, name))
}
