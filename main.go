package main

import (
	"fmt"

	"ex/exercises"
)

func main() {

	// Parse a stream of tokens into a run-length encoding.
	tokens := []string{"a", "a", "a", "b", "b", "c", "c", "d", "d", "e"}
	rle_pairs := exercises.RLE_iter(tokens)
	fmt.Printf("pairs: %v\n", rle_pairs)

	// RLE with functional-like impl
	rle_pairs = exercises.RLE_rec(tokens)
	fmt.Printf("pairs: %v\n", rle_pairs)

	// Pack consecutive duplicates of list elements into sublists.
	packed := exercises.Pack([]int{1, 1, 1, 2, 3, 3, 3, 4, 4, 4, 5})
	fmt.Printf("packed: %v\n", packed)
}
