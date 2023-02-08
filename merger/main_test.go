package main

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{3, 6, 9, 12, 15}

	expected := []int{1, 2, 3, 3, 4, 5, 6, 6, 9, 12, 15}

	c := merge(a, b)

	if len(c) != len(expected) {
		fmt.Printf("Merge result: %+v\n", c)
		t.Fatalf("Expected length %d, got %d", len(expected), len(c))
	}
	for i, v := range c {
		if v != expected[i] {
			fmt.Printf("Merge result: %+v\n", c)
			t.Fatalf("Expected %d, found %d at index %d", expected[i], v, i)
		}
	}
}

// 🟣 TODO benchmarking comparison
