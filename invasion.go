package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("We come in peace.")
	createAliens(getNumAliens())
}

func getNumAliens() int {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("missing required first parameter for number of aliens"))
	}
	fmt.Println(os.Args)
	numAliens, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	return numAliens
}

func createAliens(numAliens int) {
	fmt.Printf("Create %d aliens\n", numAliens)
}
