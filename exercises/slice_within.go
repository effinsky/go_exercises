package exercises

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
