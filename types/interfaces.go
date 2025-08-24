package types

import "github.com/roshankumar18/go-load-balancer/internal/backend"

type Strategy interface {
	NextBackend(backends []*backend.Backend) *backend.Backend
}
