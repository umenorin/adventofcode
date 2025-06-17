package main

import (
	"testing"
)

func TestPointDIrection(t *testing.T) {
	t.Run("Testing the Point move", func(t *testing.T) {
		mapSanta := MapSanta{Position{0, 0}, make(map[Position]int)}
		mapSanta.ChangePoint("v")
		want := Position{0, -1}
		if mapSanta.Position != want {
			t.Errorf("GOT %v  WANT %v\n", mapSanta.Position, want)
		}
	})
	t.Run("Testing the House that receive presents", func(t *testing.T) {
		mapSanta := MapSanta{Position{0, 0}, make(map[Position]int)}
		mapSanta.StartDelivery()
		mapSanta.ChangePoint("^v^v^v^v^v")
		want := 6

		if mapSanta.Map[Position{0, 0}] != want {
			t.Errorf("GOT %d  WANT %d\n", mapSanta.Map[Position{0, 0}], want)
		}
	})
	t.Run("Checking total of the House has visited", func(t *testing.T) {
		mapSanta := MapSanta{Position{0, 0}, make(map[Position]int)}
		mapSanta.StartDelivery()
		mapSanta.ChangePoint("^v^v^v^v^v")
		want := 2

		if mapSanta.MapSize() != want{

			t.Errorf("GOT %d  WANT %d\n", mapSanta.Map[Position{0, 0}], want)
		}
	})
}
