package algorithms

import (
	"sync/atomic"

	"github.com/roshankumar18/go-load-balancer/internal/backend"
	"github.com/roshankumar18/go-load-balancer/types"
)

type RoundRobinBalancer struct {
	current uint64
}

type StrategyType string

const (
	RoundRobin StrategyType = "round_robin"
)

func NewAlgorithm(name string) types.Strategy {

	switch StrategyType(name) {
	case RoundRobin:
		return newRoundRobinBalancer()
	default:
		return newRoundRobinBalancer()
	}
}

func newRoundRobinBalancer() *RoundRobinBalancer {
	return &RoundRobinBalancer{
		current: 0,
	}
}

func (rr *RoundRobinBalancer) NextBackend(backends []*backend.Backend) *backend.Backend {
	for i := 0; i < len(backends); i++ {
		next := int(atomic.AddUint64(&rr.current, 1) % uint64(len(backends)))
		b := backends[next]
		if b.IsAlive() {
			return b
		}
	}
	return nil
}
