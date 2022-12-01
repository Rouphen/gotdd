package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}
func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}, nil}
	handler := http.HandlerFunc(server.ServeHTTP)

	if err := http.ListenAndServe(":5008", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
