package simulation

import (
	"testing"

	. "github.com/zarazan/invasion/models"
)

func TestAliveAliens(t *testing.T) {
	createAliens(5)
	aliens[3].Destroyed = true
	if len(aliveAliens()) != 4 {
		t.Errorf("Expected 4 aliens alive")
	}
}
func TestMoveAliens(t *testing.T) {
	cities = nil
	cities = append(cities, &City{Name: "Denver", Roads: make(map[string]*City)})
	cities = append(cities, &City{Name: "Steamboat", Roads: make(map[string]*City)})

	cities[0].PaveRoad(cities[1], "west")
	createAliens(2)

	alien1Location := aliens[0].Location
	alien2Location := aliens[1].Location
	moveAliens()

	if alien1Location == aliens[0].Location {
		t.Errorf("Expected Alien 1's Location to be different")
	}
	if alien2Location == aliens[1].Location {
		t.Errorf("Expected Alien 2's Location to be different")
	}
}

func TestResolveFights(t *testing.T) {
	cities = nil
	cities = append(cities, &City{Name: "Steamboat", Roads: make(map[string]*City)})
	createAliens(2)
	resolveFights()
	if !aliens[0].Destroyed {
		t.Errorf("Expected Alien 1 to be Destroyed")
	}
	if !aliens[1].Destroyed {
		t.Errorf("Expected Alien 2 to be Destroyed")
	}
}
