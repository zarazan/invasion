package utils

import (
	"errors"
	"math/rand"
)

// GetRandomItem returns a random item from the passed in array
// An error is returned if the array is empty
func GetRandomItem[K comparable](array []K) (K, error) {
	arrayLength := len(array)
	if arrayLength < 1 {
		var result K
		return result, errors.New("cannot retrieve random item from empty array")
	}
	index := rand.Intn(arrayLength)
	return array[index], nil
}
