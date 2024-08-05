package main

import (
	"log"
	"net/http"
)

func main() {
	store:= &InMemoryPlayerStore{make(map[string]int)}
	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":8000", server))

}
