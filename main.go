package main

import (
	"net/http"

	"github.com/crdueck/interface-demo/api"
	"github.com/crdueck/interface-demo/service"
	"github.com/crdueck/interface-demo/storage"
)

func main() {
	stg := storage.New()
	svc := service.New(stg)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: api.New(svc),
	}

	srv.ListenAndServe()
}
