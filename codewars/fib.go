package codewars

func Fib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	prevPrev, prev, curr := 0, 1, 0
	for i := 2; i < n; i++ {
		curr = prev + prevPrev
		prevPrev = prev
		prev = curr
	}

	return curr
}
