package exercises

type Pair[A comparable] struct {
	Item  A
	Count int
}

// just more descriptive for the func
func NewPair[A comparable](it A) Pair[A] {
	return Pair[A]{it, 1}
}

func RLE[A comparable](in []A) []Pair[A] {
	if len(in) == 0 {
		return []Pair[A]{}
	}

	out := make([]Pair[A], 0, len(in))
	currPair := NewPair(in[0])

	for _, it := range in[1:] {
		if it == currPair.Item {
			currPair.Count++
		} else {
			out = append(out, currPair)
			currPair = NewPair(it)
		}
	}

	out = append(out, currPair)
	return out
}
