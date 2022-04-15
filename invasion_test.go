package main

import (
	"os"
	"testing"
)

func TestGetNumAliens(t *testing.T) {
	os.Args = []string{"", "5"}
	desiredResult := getNumAliens()
	if desiredResult != 5 {
		t.Errorf("Expected %d but got %d", 5, desiredResult)
	}
}
