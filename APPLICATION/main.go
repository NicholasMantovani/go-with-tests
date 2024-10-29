package main

import (
	"application/server"
	"log"
	"net/http"
)

func main() {
	srv := &server.PlayerServer{server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", srv))
}
