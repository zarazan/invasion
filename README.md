# Invasion

Invasion is a simulation of an alien invasion written in GO.

Aliens have flown to earth as a last-ditch attempt to find a habitable planet because they Destroyed their own fighting over land. Their spceship blows up in the atmosphere and they land scattered across the world in random cities.

## Installation

Close this repo to your corresponding GOPATH and run `go install .`

## Usage

`invasion [OPTION] [NUM_ALIENS]`

To run simulation with six aliens: `$ invasion 6`

Verbose logging: `$ invasion -v 6`

Specify a different worlds file: `$ invasion -world=world_2.txt 6`

* worlds files must be located in the worlds folder

## Run Tests

`go test ./...`

## Simulation Requirements

If two aliens are in the same city they fight, destroy each other, destroy the city, and destroy any adjacent roads.

Runs for a max of 10,000 moves or until all aliens have been Destroyed.

Prints to the console each time aliens fight each other.

Prints what is left of the world at the end.

## Assumptions

The aliens fight when first landing.

Aliens will always move in a direction available to them.

City Names do not contain spaces or equal signs (=)

Directions included are only "north", "south", "east" and "west". If you provide directions other than these the simulation will work but the road will only be available in one direction.

The input file could have overlapping streets such as two cities that have the same city "west" of them. Withe the exisitng functionaltiy the import function will just override the the previous street.

It is possible the city file contains no cities and the aliens do not land on any cities to start. In this situation the simulation runs 10,000 times skipping over all aiens and then completes with an empty output listing all the un-Destroyed cities. The program could be halted when no cities are present if that was the desired functionality.

## Possible Improvements

Additional error handling for the parsing of the worlds file.

An end-to-end test suite could be added that provides an input file and hooks into the console buffer to examine the output.
