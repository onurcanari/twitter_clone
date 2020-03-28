package main

import (
	"log"
	"net/http"

	Server "github.com/onurcanari/kartaca_spa/pkg/server"
)

func main() {
	r := Server.CreateServer()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", r))
}
