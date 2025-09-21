package main

import (
	"fmt"
	"net/http"
)

func main() {
	servers := []string{
		"http://server:9080",
		"http://server:9081",
		"http://server:9082",
	}

	for _, server := range servers {
		go startServer(server)
	}

	select {}
}

func startServer(url string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from %s\n", url)
	})

	port := url[len("http://server:"):]
	fmt.Printf("Starting server at %s\n", url)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Error starting server at %s: %v\n", url, err)
	}
}
