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
	// хост по которому обрашемся (роут )
	host := "0.0.0.0"

	// порт роута
	port := "9999"
	// визивае фнукцию запуска сервера
	if err := execute(host, port); err != nil {
		// если вернул ошибку остановис сервер
		os.Exit(1)
	}
}

// функция запск сервера
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
