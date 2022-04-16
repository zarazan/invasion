package main

import "testing"

func TestAliveAliens(t *testing.T) {
	createAliens(5)
	aliens[3].destroyed = true
	if len(aliveAliens()) != 4 {
		t.Errorf("Expected 4 aliens alive")
	}
}
