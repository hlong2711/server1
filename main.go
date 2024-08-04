package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string, score int) {}
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 0
}

func main() {
	server:= &PlayerServer{&InMemoryPlayerStore{make(map[string]int)}}

	log.Fatal(http.ListenAndServe(":8000", server))

}
