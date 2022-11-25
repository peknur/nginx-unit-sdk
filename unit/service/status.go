package service

import (
	"context"

	"github.com/peknur/nginx-unit-sdk/unit/status"
)

func (s *Service) Status(ctx context.Context) (status.Status, error) {
	st := status.Status{}
	return st, s.client.Get(ctx, statusPath, &st)
}
