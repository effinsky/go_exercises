package main

import (
	"fmt"

	"ex/exercises"
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

	// Rotate slice(list) right and left.
	to_rotate := []string{"a", "b", "c", "d", "e", "f", "g"}
	places := 3
	fmt.Printf("to_rotate by %d places: %v\n", places, to_rotate)

	rotated_right := ex.RotateRight(to_rotate, places)
	fmt.Printf("rotated right : %v\n", rotated_right)

	rotated_left := ex.RotateLeft(to_rotate, places)
	fmt.Printf("rotated left : %v\n", rotated_left)

	// Lotto: Draw N Different Random Numbers From the Set 1..M.
	lotto_numbers := exercises.EasyLotto(6, 50)
	fmt.Printf("lotto_numbers: %v\n", lotto_numbers)

	var s Str = "nothingness"
	fmt.Printf("s.Len(): %v\n", s.Len())

	// s = s.Push('e')
	// s = s.Push('s')
	s.PushMut('e')
	s.PushMut('s')

	fmt.Printf("s: %v\n", s)
}

type Str string

func (s Str) Len() int { return len(s) }

// returns a new Str
func (s Str) Push(ch rune) Str {
	return Str(fmt.Sprintf("%s%s", s, string(ch)))
}

// mutates Str
func (s *Str) PushMut(ch rune) {
	*s = Str(fmt.Sprintf("%s%s", *s, string(ch)))
}
