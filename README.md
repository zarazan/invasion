## Invasion

Invasion is a simulation of an alien invasion written in GO.

Aliens have flown to earth as a last-ditch attempt to find a habitable planet because they destroyed their own fighting over land. Their spceship blows up in the atmosphere and they land scattered across the world in random cities.

## Installation

`go get github.com/zarazan/invasion`

## Usage

Input: number of aliens to create

## Simulation Requirements

If two aliens are in the same city they fight, destroy each other, destroy the city, and destroy any adjacent roads.

Runs for a max pf 10,000 moves or until all aliens have been destroyed.

Prints to the console each time aliens fight each other.

Prints what is left of the world at the end.

## Tests

Run the tests: `go test . -v`

## TODO

Error handling for the parsing of the worlds file

