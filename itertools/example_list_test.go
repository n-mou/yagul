package itertools_test

import (
	"container/list"
	"fmt"

	"github.com/n-mou/yagul/itertools"
)

// A type that implements a PullIterator interface
type ListIterator struct {
	currentNode *list.Element
}

func newListIterator(l *list.List) *ListIterator {
	return &ListIterator{l.Front()}
}

func (l *ListIterator) Next() (any, bool) {
	if l.currentNode == nil {
		return nil, false
	}
	returnVal := l.currentNode.Value
	l.currentNode = l.currentNode.Next()
	return returnVal, true
}

func (l *ListIterator) Stop() {
	l.currentNode = nil
}

func ExamplePullIterator() {
	values := []string{"A", "B", "C", "D", "E"}
	l := list.New()
	for _, i := range values {
		l.PushBack(i)
	}

	iter := itertools.PullToPush(newListIterator(l))
	for i := range iter {
		fmt.Println(i)
	}
	// Output:
	// A
	// B
	// C
	// D
	// E
}
