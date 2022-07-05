package simulation

import (
	"fmt"
	"sync"

	. "github.com/zarazan/invasion/models"
	"github.com/zarazan/invasion/output"
	"github.com/zarazan/invasion/utils"
)

var (
	cities      []*City
	aliens      []*Alien
	loggingFlag bool
)

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
		if location, err := utils.GetRandomItem(cities); err == nil {
			newAlien.Location = location
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
	var waitGroup sync.WaitGroup
	for _, alien := range aliveAliens() {
		waitGroup.Add(1)
		go func(alien *Alien) {
			defer waitGroup.Done()
			moveAlien(alien)
		}(alien)
	}
	waitGroup.Wait()
}

func moveAlien(alien *Alien) {
	adjacentCities := alien.Location.AdjacentCities()
	newCity, err := utils.GetRandomItem(adjacentCities)
	if err != nil {
		printLog(fmt.Sprintf("%s cannot move because there are no adjacent cities\n", alien.Name))
		return
	}
	printLog(fmt.Sprintf("%s moving from %s to %s\n", alien.Name, alien.Location.Name, newCity.Name))
	alien.Location = newCity
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
		checkCityForFight(city, occupyingAliens)
	}
}

func checkCityForFight(city *City, aliens []*Alien) {
	numAliens := len(aliens)
	if numAliens > 1 {
		city.Destroyed = true
		for _, alien := range aliens {
			alien.Destroyed = true
		}
		output.PrintDestroyedCity(city, aliens)
	}
}

func printLog(log string) {
	if !loggingFlag {
		return
	}
	fmt.Print(log)
}
