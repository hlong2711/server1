package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{make(map[string]int)}}

	log.Fatal(http.ListenAndServe(":8000", server))

}
