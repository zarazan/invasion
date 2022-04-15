package main

import "fmt"

func PrintCities() {
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
