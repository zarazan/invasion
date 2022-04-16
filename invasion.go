package main

import (
	"errors"
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var loggingFlag bool

func main() {
	rand.Seed(time.Now().UnixNano())
	setFlags()

	numAliens, err := getNumAliensArg()
	if err != nil {
		log.Fatal(err)
	}

	readWorldFile("worlds/world_1.txt")
	runSimulation(numAliens)
	printStandingCities()
}

func setFlags() {
	flag.BoolVar(&loggingFlag, "v", false, "display verbose console logging")
	flag.Parse()
}

// Reads and parses the first command line argument
func getNumAliensArg() (int, error) {
	args := flag.Args()
	if len(args) < 1 {
		return 0, errors.New("missing required first parameter for number of aliens")
	}
	numAliens, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, err
	}
	if numAliens < 1 {
		return 0, errors.New("there must be at least one alien")
	}
	return numAliens, nil
}
