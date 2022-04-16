package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// City names cannot have any spaces
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

var oppositeDirection = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
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
