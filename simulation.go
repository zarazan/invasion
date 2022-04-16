package main

import (
	"fmt"

	"github.com/zarazan/invasion/utils"
)

var cities []*City
var aliens []*Alien

type City struct {
	name      string
	destroyed bool
	roads     map[string]*City
}

type Alien struct {
	name      string
	destroyed bool
	location  *City
}

// adjacentCities returns the un-destroyed adjacent cities that
// can be traveled to by a single road
func (c *City) adjacentCities() (ret []*City) {
	for _, city := range c.roads {
		if !city.destroyed {
			ret = append(ret, city)
		}
	}
	return
}

func runSimulation(numAliens int) {
	createAliens(numAliens)
	resolveFights() //Resolve initial fights when the aliens first land

	for i := 0; i < 10000 && len(aliveAliens()) > 0; i++ {
		moveAliens()
		resolveFights()
	}
}

func createAliens(numAliens int) {
	for i := 1; i <= numAliens; i++ {
		newAlien := &Alien{name: fmt.Sprintf("Alien %d", i)}
		aliens = append(aliens, newAlien)
		if location, err := utils.GetRandomItem(cities); err == nil {
			newAlien.location = location
		}
	}
}

func aliveAliens() (ret []*Alien) {
	for _, alien := range aliens {
		if !alien.destroyed {
			ret = append(ret, alien)
		}
	}
	return
}

// moveAliens moves all alive aliens to an available adjacent city
func moveAliens() {
	for _, alien := range aliveAliens() {
		adjacentCities := alien.location.adjacentCities()
		newCity, err := utils.GetRandomItem(adjacentCities)
		if err != nil {
			printLog(fmt.Sprintf("%s cannot move because there are no adjacent cities\n", alien.name))
			continue
		}
		printLog(fmt.Sprintf("%s moving from %s to %s\n", alien.name, alien.location.name, newCity.name))
		alien.location = newCity
	}
}

// resolveFights destroys all cities and aliens where more than 1 alien
// currenly occupies a city
func resolveFights() {
	occupationMap := make(map[*City][]*Alien)
	for _, alien := range aliveAliens() {
		if alien.location == nil {
			continue
		}
		occupationMap[alien.location] = append(occupationMap[alien.location], alien)
	}

	for city, occupyingAliens := range occupationMap {
		numAliens := len(occupyingAliens)
		if numAliens > 1 {
			city.destroyed = true
			for _, alien := range occupyingAliens {
				alien.destroyed = true
			}
			printDestroyedCity(city, occupyingAliens)
		}
	}
}
