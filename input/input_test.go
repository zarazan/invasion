package input

import (
	"fmt"
	"testing"

	. "github.com/zarazan/invasion/models"
)

func TestFindOrCreateCity(t *testing.T) {
	t.Run("when the city does not exist", func(t *testing.T) {
		cities = []*City{}
		ans := findOrCreateCity("Toronto")
		if ans.Name != "Toronto" {
			t.Errorf("expecting city Named Toronto got %s", ans.Name)
		}
		if len(cities) != 1 {
			t.Errorf("a new city was not added")
		}
	})

	t.Run("when the city does exist", func(t *testing.T) {
		cities = []*City{
			{Name: "Toronto"},
			{Name: "Lawrence"},
		}
		ans := findOrCreateCity("Toronto")
		if ans.Name != "Toronto" {
			t.Errorf("expecting city Named Toronto got %s", ans.Name)
		}
		if len(cities) != 2 {
			t.Errorf("expected the number of cities to remain unchanged")
		}
	})
}

func TestFindCityByName(t *testing.T) {
	cities = []*City{
		{Name: "Steamboat"},
		{Name: "Denver"},
		{Name: "Arvada"},
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
				t.Errorf("Expected to find city with Name %s", tt.input)
			}
			if !tt.isFound && ans != nil {
				t.Errorf("Expected not to find city with Name %s", tt.input)
			}
		})
	}
}
