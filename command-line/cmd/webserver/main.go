package main

import (
	"log"
	"net/http"
	"os"

	poker "gowithtests.com/gotdd/command-line/v2"
)

func main() {
	const dbFileName = "game.db.json"
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5008", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
