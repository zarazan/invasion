package main

import (
	"fmt"
	"testing"
)

func TestFindOrCreateCity(t *testing.T) {
	t.Run("when the city does not exist", func(t *testing.T) {
		cities = []*City{}
		ans := findOrCreateCity("Toronto")
		if ans.name != "Toronto" {
			t.Errorf("expecting city named Toronto got %s", ans.name)
		}
		if len(cities) != 1 {
			t.Errorf("a new city was not added")
		}
		teardown()
	})

	t.Run("when the city does exist", func(t *testing.T) {
		cities = []*City{
			{name: "Toronto"},
			{name: "Lawrence"},
		}
		ans := findOrCreateCity("Toronto")
		if ans.name != "Toronto" {
			t.Errorf("expecting city named Toronto got %s", ans.name)
		}
		if len(cities) != 2 {
			t.Errorf("expected the number of cities to remain unchanged")
		}
		teardown()
	})
}

func TestFindCityByName(t *testing.T) {
	cities = []*City{
		{name: "Steamboat"},
		{name: "Denver"},
		{name: "Arvada"},
	}
	var testCases = []struct {
		input   string
		isFound bool
	}{
		{"Steamboat", true},
		{"Denver", true},
		{"BadCity", false},
		{"", false},
	}

	for _, tt := range testCases {
		testName := fmt.Sprintf("input:%s, isFound:%t", tt.input, tt.isFound)
		t.Run(testName, func(t *testing.T) {

			ans := findCityByName(tt.input)
			if tt.isFound && ans == nil {
				t.Errorf("Expected to find city with name %s", tt.input)
			}
			if !tt.isFound && ans != nil {
				t.Errorf("Expected not to find city with name %s", tt.input)
			}
		})
		teardown()
	}
}
