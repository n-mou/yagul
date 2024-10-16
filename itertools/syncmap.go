package itertools

import (
	"iter"
	"sync"
)

// Returns an iterator for [sync.Map] type so it can be used in a for{} block. I
// developed it while creating the examples of this module but due to how [sync.Map]
// is implemented, the iterator would need to traverse all the map elements (which is O(N))
// twice. So I switched to a regular push style iterator and leave it as an easter egg.
func SyncMapIterator(m sync.Map) iter.Seq2[any, any] {
	return func(yield func(any, any) bool) {
		m.Range(func(key, val any) bool {
			if !yield(key, val) {
				return false
			}
			return true
		})
	}
}
