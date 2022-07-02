package simulation

import (
	"fmt"

	. "github.com/zarazan/invasion/models"
	"github.com/zarazan/invasion/output"
	"github.com/zarazan/invasion/utils"
)

var cities []*City
var aliens []*Alien
var loggingFlag bool

func RunSimulation(citiesInput []*City, numAliens int, loggingFlagInput bool) {
	loggingFlag = loggingFlagInput
	cities = citiesInput
	createAliens(numAliens)
	resolveFights() //Resolve initial fights when the aliens first land

	for i := 0; i < 10000 && len(aliveAliens()) > 0; i++ {
		moveAliens()
		resolveFights()
	}
	output.PrintStandingCities(cities)
}

func createAliens(numAliens int) {
	aliens = nil
	for i := 1; i <= numAliens; i++ {
		newAlien := Alien{Name: fmt.Sprintf("Alien %d", i)}
		aliens = append(aliens, &newAlien)
		if Location, err := utils.GetRandomItem(cities); err == nil {
			newAlien.Location = Location
		}
	}
}

func aliveAliens() (ret []*Alien) {
	for _, alien := range aliens {
		if !alien.Destroyed {
			ret = append(ret, alien)
		}
	}
	return
}

// moveAliens moves all alive aliens to an available adjacent city
func moveAliens() {
	for _, alien := range aliveAliens() {
		adjacentCities := alien.Location.AdjacentCities()
		newCity, err := utils.GetRandomItem(adjacentCities)
		if err != nil {
			printLog(fmt.Sprintf("%s cannot move because there are no adjacent cities\n", alien.Name))
			continue
		}
		printLog(fmt.Sprintf("%s moving from %s to %s\n", alien.Name, alien.Location.Name, newCity.Name))
		alien.Location = newCity
	}
}

// resolveFights destroys all cities and aliens where more than 1 alien
// currenly occupies a city
func resolveFights() {
	occupationMap := make(map[*City][]*Alien)
	for _, alien := range aliveAliens() {
		if alien.Location == nil {
			continue
		}
		occupationMap[alien.Location] = append(occupationMap[alien.Location], alien)
	}

	for city, occupyingAliens := range occupationMap {
		numAliens := len(occupyingAliens)
		if numAliens > 1 {
			city.Destroyed = true
			for _, alien := range occupyingAliens {
				alien.Destroyed = true
			}
			output.PrintDestroyedCity(city, occupyingAliens)
		}
	}
}

func printLog(log string) {
	if !loggingFlag {
		return
	}
	fmt.Print(log)
}
