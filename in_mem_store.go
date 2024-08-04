package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string, score int) {}
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 0
}
