package health

import (
	"context"
	"log"
	"time"

	"github.com/roshankumar18/go-load-balancer/internal/backend"
	"github.com/roshankumar18/go-load-balancer/internal/config"
	"github.com/roshankumar18/go-load-balancer/internal/pool"
)

type Health struct {
	serverPool   *pool.Pool
	serverConfig *config.HealthCheckConfig
}

func NewHealth(serverPool *pool.Pool, healthConfig *config.HealthCheckConfig) *Health {
	return &Health{
		serverPool:   serverPool,
		serverConfig: healthConfig,
	}
}

func (this *Health) check(backend *backend.Backend) bool {
	return backend.IsAlive()
}

func (this *Health) checkAll() bool {
	for _, backend := range this.serverPool.GetBackends() {
		alive := backend.IsBackendAlive()
		backend.SetAlive(alive)
	}

	log.Print("Total backends: ", this.serverPool.GetBackendsCount())
	log.Print("Live backends: ", this.serverPool.GetLiveBackendsCount())
	return true
}

func (this *Health) Start(ctx context.Context) {

	ticker := time.NewTicker(this.serverConfig.Interval)

	this.checkAll()

	for {
		select {
		case <-ctx.Done():
			log.Print("Health checker stopped")
			return
		case <-ticker.C:
			this.checkAll()
		}
	}
}
