package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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
	ReadWorldFile("worlds/world_1.txt")
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
	fmt.Printf("Create %d aliens\n", numAliens)
}
