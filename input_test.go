package main

import (
	"fmt"
	"testing"
)

func TestFindCityByName(t *testing.T) {
	cities = []*City{
		{name: "Steamboat"},
		{name: "Winter Park"},
		{name: "Santa Fe"},
	}
	var testCases = []struct {
		input   string
		isFound bool
	}{
		{"Steamboat", true},
		{"Santa Fe", true},
		{"Bad City", false},
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
	}
}
