package main

import (
	"log"
	"net/http"

	"github.com/leonardo404-code/proxy-server/proxy"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", proxy.Proxy)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("proxy initiated")
	log.Fatal(server.ListenAndServe())
}
