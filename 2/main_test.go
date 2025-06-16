package main

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	rectangle := Rectangle{Width: 3, Height: 4, Lenght: 2}
	got := rectangle.Area()
	want := 52

	assetsCorrectMessage(t,got,want)
}

func TestFoundSmallAreaSide(t *testing.T) {
	rectangle := Rectangle{Width: 1, Height: 1, Lenght: 10}
	got := rectangle.FoundSmallAreaSide()
	want := 1


	assetsCorrectMessage(t,got,want)
}

func TestTotalAreaNeeded(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 1, Lenght: 1}
	got := rectangle.TotalAreaNeeded()
	fmt.Println(rectangle.FoundSmallAreaSide())
	want := 43


	assetsCorrectMessage(t,got,want)
}


func assetsCorrectMessage(t testing.TB,got,want int) {
	t.Helper()

	if got != want {
		t.Errorf("GOT %d, WANT %d", got, want)
	}

}
