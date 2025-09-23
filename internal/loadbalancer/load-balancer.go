package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"

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

	log.Print(peer.GetURL())
	peer.ReverseProxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		log.Printf("[%s] %s\n", peer.GetURL().Host, e.Error())

	}

	peer.ReverseProxy = &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = peer.GetURL().Scheme
			req.URL.Host = peer.GetURL().Host
			req.Host = peer.GetURL().Host
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, e error) {
			log.Printf("[%s] %s\n", peer.GetURL().Host, e.Error())
		},
	}

	peer.ReverseProxy.ServeHTTP(w, r)

}
