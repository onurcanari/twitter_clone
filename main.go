package main

import (
	"log"
	"net/http"

	Server "github.com/onurcanari/kartaca_spa/pkg/server"
	WS "github.com/onurcanari/kartaca_spa/pkg/websocket"
)

func main() {
	hub := WS.NewHub()
	go hub.Run()
	r := Server.CreateServer(hub)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", r))
}
