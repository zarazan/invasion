package input

import (
	"bufio"
	"log"
	"os"
	"strings"

	. "github.com/zarazan/invasion/models"
)

var cities []*City

// ReadWorldFile receives the Name of the world import file and
// parses it into the cities data structure located in the simulation.go file
// City Names cannot have any spaces
// check for unsupported direction (it just won't connect up with another city)
func ReadWorldFile(fileName string) []*City {
	cities = nil

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
			city.PaveRoad(toCity, roadDirection)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return cities
}

func findOrCreateCity(Name string) *City {
	city := findCityByName(Name)
	if city != nil {
		return city
	}
	return createCity(Name)
}

func createCity(Name string) *City {
	city := &City{Name: Name, Roads: make(map[string]*City)}
	cities = append(cities, city)
	return city
}

func findCityByName(Name string) *City {
	for _, city := range cities {
		if city.Name == Name {
			return city
		}
	}
	return nil
}
