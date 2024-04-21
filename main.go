package main

import (
	"fmt"

	ds "various/data_structures"
	"various/ex"
)

func main() {
	// Parse a stream of tokens into a run-length encoding.
	tokens := []string{"a", "a", "a", "b", "b", "c", "c", "d", "d", "e"}
	rle_pairs := ex.RLE_iter(tokens)
	fmt.Printf("pairs: %v\n", rle_pairs)

	fmt.Printf("--------------------------------------------------------------\n")

	// RLE with functional-like impl
	rle_pairs = ex.RLE_rec(tokens)
	fmt.Printf("pairs: %v\n", rle_pairs)

	fmt.Printf("--------------------------------------------------------------\n")

	// Pack consecutive duplicates of list elements into sublists.
	packed := ex.Pack([]int{1, 1, 1, 2, 3, 3, 3, 4, 4, 4, 5})
	fmt.Printf("packed: %v\n", packed)

	fmt.Printf("--------------------------------------------------------------\n")
	// Slice within.
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

	fmt.Printf("--------------------------------------------------------------\n")

	// Rotate slice(list) right and left.
	to_rotate := []string{"a", "b", "c", "d", "e", "f", "g"}
	places := 3
	fmt.Printf("to_rotate by %d places: %v\n", places, to_rotate)

	rotated_right := ex.RotateRight(to_rotate, places)
	fmt.Printf("rotated right : %v\n", rotated_right)

	rotated_left := ex.RotateLeft(to_rotate, places)
	fmt.Printf("rotated left : %v\n", rotated_left)

	// Lotto: Draw N Different Random Numbers From the Set 1..M.
	lotto_numbers := ex.EasyLotto(6, 50)
	fmt.Printf("lotto_numbers: %v\n", lotto_numbers)

	fmt.Printf("--------------------------------------------------------------\n")

	// Binary search tree
	bst := ds.NewBST("alpha")
	bst.Insert("beta")
	bst.Insert("gamma")
	bst.Insert("delta")

	bst.TraversePreOrder()

	// Take another look at deletion..
	toDelete := "gamma"
	if ok := bst.Delete(toDelete); ok {
		fmt.Printf("tree traversal following %q deletion\n", toDelete)
	} else {
		fmt.Printf("Could not delete %q from tree\n", toDelete)
	}
	bst.TraversePreOrder()
}
