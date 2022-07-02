package models

type City struct {
	Name      string
	Destroyed bool
	Roads     map[string]*City
}

type Alien struct {
	Name      string
	Destroyed bool
	Location  *City
}

var oppositeDirection = map[string]string{
	"north": "south",
	"south": "north",
	"east":  "west",
	"west":  "east",
}

// adjacentCities returns the un-Destroyed adjacent cities that
// can be traveled to by a single road
func (c *City) AdjacentCities() (ret []*City) {
	for _, city := range c.Roads {
		if !city.Destroyed {
			ret = append(ret, city)
		}
	}
	return
}

// PaveRoad assigns both directions so aliens can travel back and forth
// between two connecting cities
func (c *City) PaveRoad(toCity *City, direction string) {
	c.Roads[direction] = toCity
	toCity.Roads[oppositeDirection[direction]] = c
}
