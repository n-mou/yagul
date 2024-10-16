package list

import (
	"fmt"
	"strings"
)

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List[T any] struct {
	front  *Element[T]
	back   *Element[T] // Keep track of the last node so append and pop are O(1)
	length int
}

// New returns an initialized list either empty or with some elements.
func New[T any](elements ...T) *List[T] {
	l := List[T]{nil, nil, 0}
	for _, i := range elements {
		l.PushBack(i)
	}
	return &l
}

func (l *List[T]) remove(e *Element[T]) {
	if e.list != l {
		return
	}
	if e.prev == nil {
		// e is the front of l, set the 2nd element (if any) as new front
		if e.next == nil {
			l.front = nil
		} else {
			l.front = e.next
		}
	} else {
		e.prev.next = e.next
		if e.next != nil { // If e is not the back of l
			e.next.prev = e.prev
		}
	}
	e.prev = nil
	e.next = nil

}

func (l *List[T]) pushBack(e *Element[T]) {
	// If list is empty, e is both front and back of the list.
	if l.length == 0 {
		l.front = e
		l.back = e
		return
	}
	e.prev = l.back
	l.back.next = e
	l.back = e
}

func (l *List[T]) pushFront(e *Element[T]) {
	if l.length == 0 {
		l.front = e
		l.back = e
		return
	}
	e.next = l.front
	l.front.prev = e
	l.front = e
}

// Back returns the last element of list l or nil if the list is empty.
func (l *List[T]) Back() *Element[T] {
	return l.back
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *Element[T] {
	return l.front
}

// Init initializes or clears list l.
func (l *List[T]) Init() *List[T] {
	l.front = nil
	l.back = nil
	l.length = 0
	return l
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	if mark == nil || mark.list != l {
		return nil
	}
	if mark == l.back {
		return l.PushBack(v)
	}
	newEl := &Element[T]{Value: v, list: l}
	newEl.list = l
	newEl.prev = mark
	newEl.next = mark.next
	mark.next = newEl
	mark.next.prev = newEl
	l.length++
	return newEl
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
// The mark must not be nil.
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	if mark == nil || mark.list != l {
		return nil
	}
	if mark == l.front {
		return l.PushFront(v)
	}
	return l.InsertAfter(v, mark.Prev())
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List[T]) Len() int {
	return l.length
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	if e == nil || mark == nil || e == mark || e.list != l || mark.list != l {
		return
	}
	l.remove(e)
	e.prev = mark
	e.next = mark.next
	mark.next = e
	if e.next != nil { // If mark is not the back of l
		e.next.prev = e
	}
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	if e == nil || mark == nil || e == mark || e.list != l || mark.list != l {
		return
	}
	if mark.Prev() == nil {
		// Mark is the front of l
		l.remove(e)
		l.length-- //PushFront will increment list length
		l.PushFront(e.Value)
		return
	}
	l.MoveAfter(e, mark.Prev())
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToBack(e *Element[T]) {
	if e == nil || e.list != l || e == l.back {
		return
	}
	l.remove(e)
	l.pushBack(e)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List[T]) MoveToFront(e *Element[T]) {
	if e == nil || e.list != l || e == l.front {
		return
	}
	l.remove(e)
	l.pushFront(e)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) PushBack(v T) *Element[T] {
	el := &Element[T]{Value: v, list: l}
	l.pushBack(el)
	l.length++
	return el
}

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[T]) PushBackList(other *List[T]) {
	if other == nil {
		return
	}
	for i := range other.Iterator() {
		l.PushBack(i)
	}
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List[T]) PushFront(v T) *Element[T] {
	el := &Element[T]{Value: v, list: l}
	l.pushFront(el)
	l.length++
	return el
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List[T]) PushFrontList(other *List[T]) {
	if other == nil {
		return
	}
	for i := range other.BackwardsIterator() {
		l.PushFront(i)
	}
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil.
func (l *List[T]) Remove(e *Element[T]) T {
	if e == nil || e.list != l {
		var zeroVal T
		return zeroVal
	}
	l.remove(e)
	l.length--
	e.list = nil
	return e.Value
}

// ToSlice returns a slice with a copy of the elements present in the list.
func (l List[T]) ToSlice() []T {
	out := make([]T, 0, l.Len())
	for i := range l.Iterator() {
		out = append(out, i)
	}
	return out
}

// String returns a string representation of the list and all it's members.
func (l List[T]) String() string {
	sl := make([]string, 0, l.length)
	el := l.front
	for el != nil {
		sl = append(sl, fmt.Sprintf("%v", el.Value))
		el = el.next
	}
	return "[" + strings.Join(sl, " -> ") + "]"
}
