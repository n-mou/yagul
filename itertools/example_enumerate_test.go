package itertools_test

import (
	"container/list"
	"fmt"

	"github.com/n-mou/yagul/itertools"
)

func ExampleEnumerateP() {
	testElements := []string{"A", "B", "C", "D", "E"}
	l := list.New()
	for _, i := range testElements {
		l.PushBack(i)
	}

	iter := newListIterator(l)
	enum := itertools.EnumerateP(iter)

	for k, v := range enum {
		fmt.Println(k, v)
	}
	// Output
	// 0 A
	// 1 B
	// 2 C
	// 3 D
	// 4 E
}

func ExampleEnumerate() {
	testElements := []string{"A", "B", "C", "D", "E"}
	l := list.New()
	for _, i := range testElements {
		l.PushBack(i)
	}

	iter := newListIterator(l)
	enum := itertools.Enumerate(itertools.PullToPush(iter))

	for k, v := range enum {
		fmt.Println(k, v)
	}
	// Output
	// 0 A
	// 1 B
	// 2 C
	// 3 D
	// 4 E
}
