package exercises

import (
	"math/rand"
)

// Lotto: Draw N Different Random Numbers From the Set 1..M.
func EasyLotto(n, high int) []int {
	return shuffle_with_interrupt(1, high, n)
}

// using Durstenfeld's version of the Fisher-Yates shuffle.
func shuffle_with_interrupt(low, high, n int) []int {
	// generate a range of ints of low all the way to high
	rng := make([]int, high-low+1)
	for i := range rng {
		rng[i] = low + i
	}
	// interrupt shuffling when we have n random ints..
	for i := low; i <= high && i < n; i++ {
		newIdx := rand.Intn(high-low+1) + low
		rng[i-low], rng[newIdx-low] = rng[newIdx-low], rng[i-low]
	}
	// and return a slice of that size from the range
	return rng[:n]
}
