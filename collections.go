package go_utils

import "sync"

// ConcurrentMap implements a Map function concurrently. Passed function should be threadsafe, as it will
// be called concurrently.
func ConcurrentMap[K any, V any](in []K, f func(K) V) []V {
	res := make([]V, len(in))
	var wg sync.WaitGroup
	for i := 0; i < len(in); i++ {
		idx := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			res[idx] = f(in[idx])
		}()
	}
	wg.Wait()
	return res
}

// StableFilter implements a single-threaded Filter function. Ordering of elements is guaranteed.
func StableFilter[K any](in []K, f func(K) bool) []K {
	res := make([]K, 0, len(in))
	for i := 0; i < len(in); i++ {
		if f(in[i]) {
			res = append(res, in[i])
		}
	}
	return res
}
