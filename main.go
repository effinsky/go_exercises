package main

import (
	"fmt"

	ds "various/data_structures"
	"various/ex"
)

func main() {
	// Parse a stream of tokens into a run-length encoding.
	tokens := []string{"a", "a", "a", "b", "b", "c", "c", "d", "d", "e"}
	rlePairs := ex.RLE_iter(tokens)
	fmt.Printf("pairs: %v\n", rlePairs)

	fmt.Printf("--------------------------------------------------------------\n")

	// RLE with functional-like impl
	rlePairs = ex.RLErec(tokens)
	fmt.Printf("pairs: %v\n", rlePairs)

	fmt.Printf("--------------------------------------------------------------\n")
	// Pack consecutive duplicates of list elements into sublists.
	packed := ex.Pack([]int{1, 1, 1, 2, 3, 3, 3, 4, 4, 4, 5})
	fmt.Printf("packed: %v\n", packed)

	fmt.Printf("--------------------------------------------------------------\n")
	// Slice within.
	sliceSrc := []string{
		"nothing",
		"something",
		"anything",
		"everything",
		"something_again",
		"everything_again",
		"anything_again",
	}
	from, to := 2, 5
	sliced := ex.SliceWithin(sliceSrc, from, to)
	fmt.Printf("sliced: %v\n", sliced)
	sliced = ex.SliceWithinIterative(sliceSrc, from, to) // impl super trivial
	fmt.Printf("sliced: %v\n", sliced)

	fmt.Printf("--------------------------------------------------------------\n")

	// Rotate slice(list) right and left.
	toRotate := []string{"a", "b", "c", "d", "e", "f", "g"}
	places := 3
	fmt.Printf("to_rotate by %d places: %v\n", places, toRotate)

	rotatedRight := ex.RotateRight(toRotate, places)
	fmt.Printf("rotated right : %v\n", rotatedRight)

	rotatedLeft := ex.RotateLeft(toRotate, places)
	fmt.Printf("rotated left : %v\n", rotatedLeft)

	// Lotto: Draw N Different Random Numbers From the Set 1..M.
	lottoNumbers := ex.EasyLotto(6, 50)
	fmt.Printf("lotto_numbers: %v\n", lottoNumbers)

	fmt.Printf("--------------------------------------------------------------\n")

	// Binary search tree
	bst := ds.NewBST("alpha")

	sortedInsertSource := []string{"beta", "gamma", "delta"}
	// This bulk insertion using a binary recursion algo is to ensure the tree
	// is balanced. But it assumes the input is sorted. If it is not, it will
	// run, but the resulting tree will not be balanced as intended.
	bst.InsertFromSorted(sortedInsertSource)

	bst.Insert("delta")
	bst.Insert("gamma")
	bst.Insert("beta")

	bst.TraversePreOrder()

	// Take another look at deletion..
	// toDelete := "gamma"
	// if ok := bst.Delete(toDelete); ok {
	// 	fmt.Printf("tree traversal following %q deletion\n", toDelete)
	// } else {
	// 	fmt.Printf("Could not delete %q from tree\n", toDelete)
	// }
	// bst.TraversePreOrder()
}
