package main

import (
	"testing"
)

func teardown() {
	cities = []*City{}
	aliens = []*Alien{}
}

func TestAliveAliens(t *testing.T) {
	createAliens(5)
	aliens[3].destroyed = true
	if len(aliveAliens()) != 4 {
		t.Errorf("Expected 4 aliens alive")
	}
	teardown()
}
func TestMoveAliens(t *testing.T) {
	createCity("Denver")
	createCity("Steamboat")
	paveRoad(cities[0], cities[1], "west")
	createAliens(2)

	alien1Location := aliens[0].location
	alien2Location := aliens[1].location
	moveAliens()

	if alien1Location == aliens[0].location {
		t.Errorf("Expected Alien 1's location to be different")
	}
	if alien2Location == aliens[1].location {
		t.Errorf("Expected Alien 2's location to be different")
	}
	teardown()
}

func TestResolveFights(t *testing.T) {
	createCity("Steamboat")
	createAliens(2)
	resolveFights()
	if !aliens[0].destroyed {
		t.Errorf("Expected Alien 1 to be destroyed")
	}
	if !aliens[1].destroyed {
		t.Errorf("Expected Alien 2 to be destroyed")
	}
	teardown()
}
