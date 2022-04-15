package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cities []*City

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

func main() {
	fmt.Println("We come in peace.")
	numAliens, err := getNumAliens()
	if err != nil {
		log.Fatal(err)
	}
	createAliens(numAliens)
	readWorldFile("worlds/world_1.txt")
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
	fmt.Printf("Create %d aliens\n", numAliens)
}

func readWorldFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		cityName := words[0]

		city := findOrCreateCity(cityName)

		for _, word := range words[1:] {
			road := strings.Split(word, "=")
			roadDirection := road[0]
			roadCityName := road[1]
			toCity := findOrCreateCity(roadCityName)
			paveRoad(city, toCity, roadDirection)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func paveRoad(fromCity *City, toCity *City, direction string) {
	fromCity.roads[direction] = toCity
	toCity.roads[oppositeDirection[direction]] = fromCity
}

func findOrCreateCity(name string) *City {
	city := findCityByName(name)
	if city != nil {
		return city
	}
	city = &City{name: name, roads: make(map[string]*City)}
	cities = append(cities, city)
	return city
}

func findCityByName(name string) *City {
	for _, city := range cities {
		if city.name == name {
			return city
		}
	}
	return nil
}
