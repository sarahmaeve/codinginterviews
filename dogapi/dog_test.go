package main

import "testing"

func TestCheckBreed(t *testing.T) {
	m := getBreedList()
	bt := map[string]bool{
		"spaniel":         true,
		"Cocker Spaniel":  true,
		"giraffe":         false,
		"manx":            false,
		"pug":             true,
		"border collie":   true,
		"German Shepherd": true,
	}

	for brd, expect := range bt {
		if checkBreedList(m, brd) != expect {
			t.Errorf("expected %v to be %v, got %v", brd, expect, !expect)
		}
	}
}

func TestTopSubbreeds(t *testing.T) {
	m := getBreedList()
	top := topSubbreeds(m, 1)
	if top[0].Breed != "terrier" || top[0].Value != 20 {
		t.Errorf("did not get terrier: 20")
	}
}
