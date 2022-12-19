package main

import (
	"fmt"
	"log"
	"os"

	poker "gowithtests.com/gotdd/time"
)

const dbFileName = "game.db.json"

func main() {

	store, err := poker.FileSystemStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
