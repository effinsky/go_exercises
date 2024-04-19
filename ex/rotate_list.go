package ex

func RotateRight[A any](in []A, n int) []A {
	o := getOffset(in, n)
	return append(in[o+1:], in[:o+1]...)
}

func RotateLeft[A any](in []A, n int) []A {
	o := getOffset(in, n)

	tmp := make([]A, len(in[:o]))
	copy(tmp, in[:o])
	copy(in, in[o:])
	copy(in[len(in)-o:], tmp)

	return in
}

// Ensure n is within slice bounds
func getOffset[A any](in []A, base int) int {
	return base % len(in)
}
