package main

import (
	poker "application"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {

	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}
	defer close()

	srv := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", srv); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
