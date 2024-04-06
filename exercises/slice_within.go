package exercises

// SliceWithinIterative here creates a new list in memory, though the elements,
// if they are reference types, are not cloned, but shared. We can go all sorts
// of ways with this. Iterative is pretty trivial here compared to the
// functional aproach below. If we wanted full immutability, we'd have to do
// deep cloning on all the values in the input list. and that would not be so
// trivial anymore.. and we cannot really otherwise guarantee immutability of
// values of custom types within the language..
func SliceWithinIterative[A any](in []A, from, to int) []A {
	// not using Go slice syntax on purpose here
	out := make([]A, 0, to-from+1)
	for i := from; i <= to; i++ {
		out = append(out, in[i])
	}
	return out
}

// bonus twisted functional-style impl
func SliceWithin[A any](in []A, i, j int) []A {
	if len(in) == 0 {
		return []A{}
	}

	_, in_sans_dropped := foldUntil(
		func(_ []A, _ A) []A { return []A{} },
		[]A{},
		i,
		in,
	)
	taken, _ := foldUntil(
		func(acc []A, h A) []A { return append(acc, h) },
		[]A{},
		j-i+1,
		in_sans_dropped,
	)
	return taken
}

func foldUntil[A any, B any](
	f func(acc A, it B) A,
	acc A,
	n int,
	list []B,
) (A, []B) {
	if len(list) == 0 {
		return acc, []B{}
	}

	if n == 0 {
		return acc, list
	}

	h := list[0] // we know we have this since len > 0
	t := list[1:]
	acc = f(acc, h)
	return foldUntil(f, acc, n-1, t)
}
