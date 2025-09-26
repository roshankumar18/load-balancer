package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/roshankumar18/go-load-balancer/internal/algorithms"
	"github.com/roshankumar18/go-load-balancer/internal/config"
	"github.com/roshankumar18/go-load-balancer/internal/health"
	"github.com/roshankumar18/go-load-balancer/internal/loadbalancer"
	"github.com/roshankumar18/go-load-balancer/internal/metrics"
	"github.com/roshankumar18/go-load-balancer/internal/pool"
)

func main() {
	metrics.Init()

	config, err := config.Load("./configs/config.yaml")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	pool := pool.NewPool()
	pool.SetStrategy(algorithms.NewAlgorithm(config.LoadBalancer.Algorithm))

	log.Print(config)
	for _, backendCfg := range config.Backends {
		pool.AddBackendFromURL(backendCfg.URL)
	}

	health := health.NewHealth(pool, &config.Health)

	mux := http.NewServeMux()
	mux.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving Prometheus metrics")
		promhttp.Handler().ServeHTTP(w, r)
	}))
	mux.Handle("/", loadbalancer.NewLoadBalancer(pool))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go health.Start(ctx)

	go func() {
		fmt.Println("Starting server", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server")

	shutCtx, shutCancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer shutCancel()

	if err := server.Shutdown(shutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
