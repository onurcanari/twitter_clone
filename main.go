package main

import (
	"log"
	"net/http"

	DB "github.com/onurcanari/kartaca_spa/pkg/db"
	Server "github.com/onurcanari/kartaca_spa/pkg/server"
)

func main() {
	DB.GetUser("admin")
	r := Server.CreateServer()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", r))
}
