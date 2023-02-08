package main

import "sort"

func merge_firstIteration(a []int, b []int) []int {
	c := append(a, b...)
	sort.Ints(c)

	return c
}

func merge(a []int, b []int) []int {
	c := make([]int, 0)
	lastB := 0
	for _, ai := range a {
		for j := lastB; j < len(b); j++ {
			if b[j] <= ai {
				c = append(c, b[j])
			} else {
				lastB = j
				break
			}
		}
		c = append(c, ai)
	}

	if len(c) != len(a)+len(b) {
		c = append(c, b[lastB:]...)
	}
	return c
}
