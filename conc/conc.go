package conc

import (
	"sync"

	"golang.org/x/exp/constraints"
)

func UseConc[T constraints.Ordered](t1, t2 T) chan T {
	var wg sync.WaitGroup
	wg.Add(2)

	resChan := make(chan T)

	for _, p := range []T{t1, t2} {
		go func(p T) {
			// Chan will block until a single value from it is read.
			resChan <- p
			// This will decrement wg counter by 1.
			wg.Done()
		}(p)
	}
	go func() {
		// This will unblock once the counter on wg is zero.
		wg.Wait()
		close(resChan)
	}()

	return resChan
}
