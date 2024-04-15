package exercises

import (
	"math/rand"
)

// Lotto: Draw N Different Random Numbers From the Set 1..M.
func EasyLotto(n, high int) []int {
	// definitely not the most efficient. but not too bad either!
	return shuffle(1, high)[:n]
}

// using Durstenfeld's version of the Fisher-Yates shuffle.
func shuffle(low, high int) []int {
	rng := make([]int, 0, high-low+1)

	for i := low; i <= high; i++ {
		rng = append(rng, i)
	}

	for i := low; i < high; i++ {
		new_idx := rand.Intn(high - low)
		last_untouched_idx := high - i
		rng[new_idx], rng[last_untouched_idx] = rng[last_untouched_idx], rng[new_idx]
	}

	return rng
}
