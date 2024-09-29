package main

import (
	"log"
	"net/http"

	poker "github.com/robinkb/learn-go-with-tests/build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem opening filesystem player store: %v", err)
	}
	defer closeFunc()

	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
