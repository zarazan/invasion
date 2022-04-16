package main

import (
	"fmt"
	"strings"
)

func printStandingCities() {
	for _, city := range cities {
		if city.destroyed {
			continue
		}
		output := city.name
		for direction, adjacentCity := range city.roads {
			if !adjacentCity.destroyed {
				output = output + fmt.Sprintf(" %s=%s", direction, adjacentCity.name)
			}
		}
		fmt.Println(output)
	}
}

func printDestroyedCity(city *City, aliens []*Alien) {
	fmt.Printf("%s has been destroyed by ", city.name)
	alienNames := make([]string, 0)
	for _, alien := range aliens {
		alienNames = append(alienNames, alien.name)
	}
	last := alienNames[0]
	alienNames = alienNames[1:]
	alienList := strings.Join(alienNames, ", ") + " and " + last
	fmt.Println(alienList)
}
