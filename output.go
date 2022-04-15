package main

import "fmt"

func PrintCities() {
	for _, city := range cities {
		output := city.name
		for direction, city := range city.roads {
			output = output + fmt.Sprintf(" %s=%s", direction, city.name)
		}
		fmt.Println(output)
	}
}
