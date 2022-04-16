package utils

import "testing"

func TestGetRandomItem(t *testing.T) {
	array := []int{1, 2}
	ans, err := GetRandomItem(array)
	if err != nil {
		t.Errorf("Expected no error got error %d", err)
	}
	if ans != 1 && ans != 2 {
		t.Errorf("Expected 1 or 2 but got %d", ans)
	}
}
