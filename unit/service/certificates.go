package service

import (
	"context"
	"path"

	"github.com/peknur/nginx-unit-sdk/unit/certificate"
)

func (s *Service) Certificates(ctx context.Context) (certificate.Certificates, error) {
	c := certificate.Certificates{}
	return c, s.client.Get(ctx, certificatesPath, &c)
}

func (s *Service) CreateCertificate(ctx context.Context, name string, bundle []byte) error {
	return s.client.PutBinary(ctx, path.Join(certificatesPath, name), bundle)
}

func (s *Service) DeleteCertificate(ctx context.Context, name string) error {
	return s.client.Delete(ctx, path.Join(certificatesPath, name))
}
