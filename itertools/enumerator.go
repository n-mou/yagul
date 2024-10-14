package itertools

import "iter"

// Enumerate receives a single value iterator and returns a key-pair value
// iterator, being the key a count value that starts with 0. Each time the
// iterator it's called it returns that count and the value that the original
// iterator would have returned.
func Enumerate[T any](p iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		count := 0
		for i := range p {
			if !yield(count, i) {
				return
			}
			count++
		}
	}
}

// EnumerateP does the same as Enumerate but receives a [PullIterator] instead
// of an [iter.Seq] (it just calls [PullToPush] in the background).
func EnumerateP[T any](p PullIterator[T]) iter.Seq2[int, T] {
	return Enumerate(PullToPush(p))
}
