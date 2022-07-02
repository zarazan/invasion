package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/zarazan/invasion/input"
	"github.com/zarazan/invasion/simulation"
)

var loggingFlag bool
var fileNameFlag string

func main() {
	rand.Seed(time.Now().UnixNano())
	setFlags()

	numAliens, err := getNumAliensArg()
	if err != nil {
		log.Fatal(err)
	}

	cities := input.ReadWorldFile(fmt.Sprintf("worlds/%s", fileNameFlag))
	simulation.RunSimulation(cities, numAliens, loggingFlag)
}

func setFlags() {
	flag.BoolVar(&loggingFlag, "v", false, "display verbose console logging")
	flag.StringVar(&fileNameFlag, "world", "world_1.txt", "specify the world import file to use")
	flag.Parse()
}

// Reads and parses the first command line argument
func getNumAliensArg() (int, error) {
	if len(os.Args) < 1 {
		return 0, errors.New("missing required integer parameter for number of aliens")
	}
	numAliensArg := os.Args[len(os.Args)-1]
	numAliens, err := strconv.Atoi(numAliensArg)
	if err != nil {
		return 0, err
	}
	if numAliens < 1 {
		return 0, errors.New("there must be at least one alien")
	}
	return numAliens, nil
}
