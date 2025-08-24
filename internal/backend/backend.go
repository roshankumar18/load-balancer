package backend

import (
	"net"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type Backend struct {
	alive        bool
	url          *url.URL
	ReverseProxy *httputil.ReverseProxy
	connections  int
	mux          sync.RWMutex
}

func NewBackend(url *url.URL) *Backend {
	return &Backend{
		url:          url,
		ReverseProxy: httputil.NewSingleHostReverseProxy(url),
		alive:        true,
		connections:  0,
	}
}

func (this *Backend) SetAlive(alive bool) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.alive = alive
}

func (this *Backend) IsAlive() bool {
	this.mux.RLock()
	defer this.mux.RUnlock()
	return this.alive
}

func (this *Backend) AddConnection() {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.connections++
}

func (this *Backend) RemoveConnection() {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.connections--
}

func (this *Backend) GetConnections() int {
	this.mux.RLock()
	defer this.mux.RUnlock()
	return this.connections
}

func (this *Backend) GetURL() *url.URL {
	return this.url
}

func (this *Backend) IsBackendAlive() bool {

	conn, err := net.DialTimeout("tcp", this.url.Host, 2*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
