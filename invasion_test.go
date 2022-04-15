package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetNumAliens(t *testing.T) {
	var testCases = []struct {
		input    string
		want     int
		hasError bool
	}{
		{"5", 5, false},
		{"967", 967, false},
		{"-1", 0, true},
		{"letters", 0, true},
		{"", 0, true},
	}

	for _, tt := range testCases {
		testName := fmt.Sprintf("input:%s, expectedOutput:%d, error:%t", tt.input, tt.want, tt.hasError)
		t.Run(testName, func(t *testing.T) {

			os.Args = []string{"", tt.input}
			ans, err := getNumAliens()
			if !tt.hasError && err != nil {
				t.Errorf("Expected no error got error %d", err)
			}
			if tt.hasError && err == nil {
				t.Errorf("Expected error got %d", err)
			}
			if ans != tt.want {
				t.Errorf("Expected %d but got %d", tt.want, ans)
			}
		})
	}
}

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
