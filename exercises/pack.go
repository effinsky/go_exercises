package exercises

type Subl[A any] []A

func NewSubl[A any](items ...A) Subl[A] {
	return Subl[A](items)
}

func Pack[A comparable](in []A) []Subl[A] {
	if len(in) == 0 {
		return []Subl[A]{}
	}

	out := make([]Subl[A], 0, len(in))
	subl := NewSubl(in[0])

	for _, v := range in[1:] {
		if v == subl[0] {
			subl = append(subl, v)
		} else {
			out = append(out, subl)
			subl = NewSubl(v)
		}
	}

	out = append(out, subl)
	return out
}
