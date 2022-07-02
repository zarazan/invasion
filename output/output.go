package output

import (
	"fmt"
	"strings"

	. "github.com/zarazan/invasion/models"
)

func PrintStandingCities(cities []*City) {
	for _, city := range cities {
		if city.Destroyed {
			continue
		}
		output := city.Name
		for direction, adjacentCity := range city.Roads {
			if !adjacentCity.Destroyed {
				output = output + fmt.Sprintf(" %s=%s", direction, adjacentCity.Name)
			}
		}
		fmt.Println(output)
	}
}

func PrintDestroyedCity(city *City, aliens []*Alien) {
	fmt.Printf("%s has been Destroyed by ", city.Name)
	alienNames := make([]string, 0)
	for _, alien := range aliens {
		alienNames = append(alienNames, alien.Name)
	}
	last := alienNames[0]
	alienNames = alienNames[1:]
	alienList := strings.Join(alienNames, ", ") + " and " + last
	fmt.Println(alienList)
}
