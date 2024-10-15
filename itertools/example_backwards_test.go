package itertools_test

import (
	"container/list"
	"fmt"

	"github.com/n-mou/yagul/itertools"
)

type ListBackwardsIterator struct {
	currentNode *list.Element
}

func NewListBackwardsIterator(l *list.List) *ListBackwardsIterator {
	return &ListBackwardsIterator{l.Back()}
}

func (l *ListBackwardsIterator) Next() (any, bool) {
	if l.currentNode == nil {
		return nil, false
	}
	returnVal := l.currentNode.Value
	l.currentNode = l.currentNode.Prev()
	return returnVal, true
}

func (l *ListBackwardsIterator) Stop() {}

func ExamplePullIterator_backwards() {
	values := []string{"A", "B", "C", "D", "E"}
	l := list.New()
	for _, i := range values {
		l.PushBack(i)
	}

	iter := itertools.PullToPush(NewListBackwardsIterator(l))
	for i := range iter {
		fmt.Println(i)
	}
	// Output:
	// E
	// D
	// C
	// B
	// A
}
