package list

import "fmt"

// Element is an element of a linked list.
type Element[T any] struct {
	Value T
	list  *List[T]
	next  *Element[T]
	prev  *Element[T]
}

// Next returns the next list element or nil.
func (e *Element[T]) Next() *Element[T] {
	return e.next
}

// Prev returns the previous list element or nil.
func (e *Element[T]) Prev() *Element[T] {
	return e.prev
}

// String returns a string representation of the list node
func (e *Element[T]) String() string {
	var pVal, nVal T
	if e.prev != nil {
		pVal = e.prev.Value
	}
	if e.next != nil {
		nVal = e.next.Value
	}
	return fmt.Sprintf("%v <- [%v] -> %v", pVal, e.Value, nVal)
}
