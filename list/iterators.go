package list

import (
	"iter"
	"yagul/itertools"
)

type listIterator[T any] struct {
	el *Element[T]
}
type listBackwardsIterator[T any] struct {
	el *Element[T]
}

func (i *listIterator[T]) Next() (T, bool) {
	var zeroVal T
	if i.el == nil {
		return zeroVal, false
	}
	returnVal := i.el.Value
	i.el = i.el.Next()
	return returnVal, true
}

func (i *listIterator[T]) Stop() {
}

func (i *listBackwardsIterator[T]) Next() (T, bool) {
	var zeroVal T
	if i.el == nil {
		return zeroVal, false
	}
	returnVal := i.el.Value
	i.el = i.el.Prev()
	return returnVal, true
}

func (i *listBackwardsIterator[T]) Stop() {
}

// Iterator returns an iterator that traverses all list elements from first to last
// that can be used in the for loop.
func (l List[T]) Iterator() iter.Seq[T] {
	i := listIterator[T]{l.front}
	return itertools.PullToPush(&i)
}

// BackwardsIterator returns an iterator that traverses all list elements from last
// to first that can be used in the for loop.
func (l List[T]) BackwardsIterator() iter.Seq[T] {
	i := listBackwardsIterator[T]{l.back}
	return itertools.PullToPush(&i)
}
