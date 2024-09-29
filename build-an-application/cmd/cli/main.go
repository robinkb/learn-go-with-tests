package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/robinkb/learn-go-with-tests/build-an-application"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's player poker!")
	fmt.Println("Type '{Name} wins' to record a win.")

	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem opening filesystem player store: %v", err)
	}
	defer closeFunc()

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
