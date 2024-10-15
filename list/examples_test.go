package list_test

import (
	"fmt"

	"github.com/n-mou/yagul/list"
)

func ExampleList_Iterator() {
	l := list.New("A", "B", "C", "D")
	for i := range l.Iterator() {
		fmt.Println(i)
	}
	// Output:
	// A
	// B
	// C
	// D
}

func ExampleList_BackwardsIterator() {
	l := list.New("A", "B", "C", "D")
	for i := range l.BackwardsIterator() {
		fmt.Println(i)
	}
	// Output:
	// D
	// C
	// B
	// A
}
