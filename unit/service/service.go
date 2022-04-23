package service

import (
	"github.com/peknur/nginx-unit-sdk/unit"
)

type Service struct {
	client unit.Client
}

var _ unit.Service = (*Service)(nil)

func New(client unit.Client) Service {
	return Service{client: client}
}
