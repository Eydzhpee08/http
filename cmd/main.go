package main

import (
	// "log"
	"net"
	"net/http"
	"os"

	"github.com/Eydzhpee08/http/cmd/app"
	"github.com/Eydzhpee08/http/pkg/banners"

)
func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannersSvc := banners.NewService()
	server := app.NewServer(mux, bannersSvc)
	server.Init()

	srv := http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}

	return srv.ListenAndServe()
}
