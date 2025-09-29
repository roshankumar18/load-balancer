package health

import (
	"context"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/roshankumar18/go-load-balancer/internal/backend"
	"github.com/roshankumar18/go-load-balancer/internal/config"
	"github.com/roshankumar18/go-load-balancer/internal/metrics"
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
		value := 0
		alive := backend.IsBackendAlive()
		backend.SetAlive(alive)

		if alive {
			value = 1
		}
		metrics.BackendAlive.With(prometheus.Labels{"backend": backend.GetURL().Host}).Set(float64(value))
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
