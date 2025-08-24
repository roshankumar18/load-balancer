package pool

import (
	"net/url"
	"sync"

	"github.com/roshankumar18/go-load-balancer/internal/backend"
	"github.com/roshankumar18/go-load-balancer/types"
)

type Pool struct {
	backends []*backend.Backend
	mux      sync.RWMutex
	strategy types.Strategy
}

func NewPool() *Pool {
	return &Pool{
		backends: make([]*backend.Backend, 0),
	}
}

func (this *Pool) AddBackend(backend *backend.Backend) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.backends = append(this.backends, backend)
}

func (this *Pool) AddBackendFromURL(serverUrl string) error {
	u, err := url.Parse(serverUrl)
	if err != nil {
		return err
	}
	this.AddBackend(backend.NewBackend(u))
	return nil

}

func (this *Pool) SetStrategy(strategy types.Strategy) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.strategy = strategy
}

func (this *Pool) GetBackends() []*backend.Backend {
	return this.backends
}

func (this *Pool) GetStrategy() types.Strategy {
	return this.strategy
}

func (this *Pool) GetBackend() *backend.Backend {
	return this.strategy.NextBackend(this.backends)
}

func (this *Pool) GetBackendsCount() int {
	return len(this.backends)
}

func (this *Pool) GetLiveBackendsCount() int {
	count := 0
	for _, backend := range this.backends {
		if backend.IsAlive() {
			count++
		}
	}
	return count
}
