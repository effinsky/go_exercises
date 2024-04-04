package main

import (
	"fmt"

	ex "ex/exercises"
)

func main() {
	// Parse a stream of tokens into a run-length encoding.
	tokens := []string{"a", "a", "a", "b", "b", "c", "c", "d", "d", "e"}
	rle_pairs := ex.RLE_iter(tokens)
	fmt.Printf("pairs: %v\n", rle_pairs)

	// RLE with functional-like impl
	rle_pairs = ex.RLE_rec(tokens)
	fmt.Printf("pairs: %v\n", rle_pairs)

	// Pack consecutive duplicates of list elements into sublists.
	packed := ex.Pack([]int{1, 1, 1, 2, 3, 3, 3, 4, 4, 4, 5})
	fmt.Printf("packed: %v\n", packed)

	slice_src := []string{
		"nothing",
		"something",
		"anything",
		"everything",
		"something_again",
		"everything_again",
		"anything_again",
	}
	from, to := 2, 5
	sliced := ex.SliceWithin(slice_src, from, to)
	fmt.Printf("sliced: %v\n", sliced)
	sliced = ex.SliceWithinIterative(slice_src, from, to) // impl super trivial
	fmt.Printf("sliced: %v\n", sliced)
}
