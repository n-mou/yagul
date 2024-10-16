package itertools_test

import (
	"fmt"
	"sync"

	"github.com/n-mou/yagul/itertools"
)

func ExampleSyncMapIterator() {
	m := sync.Map{}
	m.Store("Key 1", "Val 1")
	m.Store("Key 2", "Val 2")
	m.Store("Key 3", "Val 3")

	iter := itertools.SyncMapIterator(m)
	for k, v := range iter {
		fmt.Println(k, v)
	}
	// There's no specified output beause the order
	// of the map keys is not guaranteed.
}
