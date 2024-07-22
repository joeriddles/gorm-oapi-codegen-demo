package main

import (
	"log"
	"net/http"

	"github.com/joeriddles/gorm-oapi-codegen-demo/api"
)

func main() {
	server := api.NewServer()

	r := http.NewServeMux()

	si := api.NewStrictHandler(server, []api.StrictMiddlewareFunc{})
	h := api.HandlerFromMux(si, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
