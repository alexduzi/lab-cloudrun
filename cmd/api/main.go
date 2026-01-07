package main

import (
	"log"
	"os"

	h "github.com/alexduzi/labcloudrun/internal/http"
)

func main() {
	srv := h.NewHttpHandler("8080")
	defer func() {
		log.Default().Printf("server up and running at %s port", srv.Addr)
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
