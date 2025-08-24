package loadbalancer

import (
	"log"
	"net/http"

	"github.com/roshankumar18/go-load-balancer/internal/backend"
	"github.com/roshankumar18/go-load-balancer/internal/pool"
)

type LoadBalancer struct {
	pool *pool.Pool
}

func NewLoadBalancer(pool *pool.Pool) *LoadBalancer {
	return &LoadBalancer{
		pool: pool,
	}
}

func (this *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	peer := this.pool.GetBackend()
	if peer != nil {
		this.proxyRequest(w, r, peer)
		return
	}

	http.Error(w, "No backend available", http.StatusServiceUnavailable)

}

func (lb *LoadBalancer) proxyRequest(w http.ResponseWriter, r *http.Request, peer *backend.Backend) {
	peer.AddConnection()
	defer peer.RemoveConnection()

	peer.ReverseProxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		log.Printf("[%s] %s\n", peer.GetURL().Host, e.Error())

		// retries := utils.GetRetryFromContext(request)
		// if retries < lb.config.MaxRetries {
		// 	// Retry with delay
		// 	select {
		// 	case <-time.After(lb.config.Delay):
		// 		ctx := context.WithValue(request.Context(), utils.Retry, retries+1)
		// 		peer.ReverseProxy.ServeHTTP(writer, request.WithContext(ctx))
		// 	}
		// 	return
		// }

		// // After max retries, mark this backend as down
		// lb.serverPool.MarkBackendStatus(peer.GetURL(), false)

		// // Try with a different backend
		// attempts := utils.GetAttemptsFromContext(request)
		// log.Printf("%s(%s) Attempting retry %d\n", request.RemoteAddr, request.URL.Path, attempts)

		// ctx := context.WithValue(request.Context(), utils.Attempts, attempts+1)
		// lb.ServeHTTP(writer, request.WithContext(ctx))
	}

	peer.ReverseProxy.ServeHTTP(w, r)
}
