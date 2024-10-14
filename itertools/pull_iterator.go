package itertools

import "iter"

type PullIterator[V any] interface {
	Next() (V, bool)
	Stop()
}

type PullIterator2[K any, V any] interface {
	Next() (K, V, bool)
	Stop()
}

// PullToPush transforms a [PullIterator] into an [iter.Seq]
// iterator that can be used in a "for v := range" block.
func PullToPush[V any](i PullIterator[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		val, ok := i.Next()
		for ok {
			if !yield(val) {
				i.Stop()
				return
			}
			val, ok = i.Next()
		}
		i.Stop()
	}
}

// PullToPush2 transforms a [PullIterator2] into a [iter.Seq2] iterator that can
// be used in a "for v := range" block.
func PullToPush2[K any, V any](i PullIterator2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		k, v, ok := i.Next()
		for ok {
			if !yield(k, v) {
				i.Stop()
				return
			}
			k, v, ok = i.Next()
		}
		i.Stop()
	}
}
