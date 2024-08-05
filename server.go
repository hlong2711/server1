package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string, score int)
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {

	s:= &PlayerServer{store: store}
	router := http.NewServeMux()

	router.Handle("/players/", http.HandlerFunc(s.playersHandler))

	router.Handle("/leagues", http.HandlerFunc(s.leagueHandler))

	s.Handler = router

	return s
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.postScore(w, r)

	case http.MethodGet:
		p.getScore(w, r)
	}
}

func (s *PlayerServer) getScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := s.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) postScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	s.store.RecordWin(player, 11)
	w.WriteHeader(http.StatusAccepted)
}
