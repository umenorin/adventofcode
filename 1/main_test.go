package main

import "testing"


func TestParenthesisFunction(t *testing.T) {
	list := []uint8{40,41,41,40,40}

	got,err := ParenthesisCountFloor(list)

	if err != nil {
		t.Errorf("A error has ocurred, %v",err)
	}
	want :=1

	if got != want {
		t.Errorf("GOT %d, WANT %d",got,want)
	}
}

func TestCheckWhenBackTheFistFloor(t *testing.T) {
	list := []uint8{40,41,40,41,41}

	got := CheckWhenBackTheFistFloor(list)


	want := 5

	if got != want {
		t.Errorf("GOT %d, WANT %d",got,want)
	}

	
}
