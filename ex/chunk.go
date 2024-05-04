package ex

func Chunk[A any](in []A, n int) [][]A {
	out := make([][]A, 0, len(in)/n)

	for i := 0; i < len(in); i += n {
		out = append(out, in[i:i+n])
	}

	return out
}
