package iter

type (
	Seq0           func(func() bool)
	Seq[V any]     func(func(V) bool)
	Seq2[K, V any] func(func(K, V) bool)
)

func PickOnCond[E any](s []E, pred func(in int) bool) Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := 0; i < len(s); i++ {
			if pred(i) && !yield(i, s[i]) {
				return
			}
		}
	}
}
