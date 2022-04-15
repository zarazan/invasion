package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/zarazan/invasion/utils"
)

var cities []*City
var aliens []*Alien

var oppositeDirection = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

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

func (c *City) adjacentCities() (ret []*City) {
	for _, city := range c.roads {
		if !city.destroyed {
			ret = append(ret, city)
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("We come in peace.")
	numAliens, err := getNumAliens()
	if err != nil {
		log.Fatal(err)
	}
	ReadWorldFile("worlds/world_1.txt")
	createAliens(numAliens)
	resolveFights()

	for i := 0; i < 10000; i++ {
		moveAliens()
		resolveFights()
	}

	PrintCities()
}

// Reads and parses the first command line argument
func getNumAliens() (int, error) {
	if len(os.Args) < 2 {
		return 0, errors.New("missing required first parameter for number of aliens")
	}
	numAliens, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return 0, err
	}
	if numAliens < 1 {
		return 0, errors.New("there must be at least one alien")
	}
	return numAliens, nil
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

func moveAliens() {
	for _, alien := range aliens {
		if alien.destroyed {
			continue
		}
		adjacentCities := alien.location.adjacentCities()
		newCity, err := utils.GetRandomItem(adjacentCities)
		if err != nil {
			fmt.Printf("%s cannot move because there are no adjacent cities\n", alien.name)
			continue
		}
		fmt.Printf("%s moving from %s to %s\n", alien.name, alien.location.name, newCity.name)
		alien.location = newCity
	}
}

func resolveFights() {
	occupationMap := make(map[*City][]*Alien)
	for _, alien := range aliens {
		if alien.destroyed || alien.location == nil {
			continue
		}
		occupationMap[alien.location] = append(occupationMap[alien.location], alien)
	}

	for city, occupyingAliens := range occupationMap {
		if len(occupyingAliens) > 1 {
			city.destroyed = true
			fmt.Printf("%s has been destroyed\n", city.name)
			for _, alien := range occupyingAliens {
				alien.destroyed = true
				fmt.Printf("%s has been killed\n", alien.name)
			}
		}
	}
}
