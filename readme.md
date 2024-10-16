# YAGUL

<a href="https://pkg.go.dev/github.com/n-mou/yagul"><img src="https://pkg.go.dev/badge/github.com/n-mou/yagul.svg" alt="Go Reference"></a>

YAGUL (Yet Another Go-Utils Library) is my personal recopilation of helper and quality of life improvement modules that I regularly use in my go projects and I made it public because other devs may also find them useful. 


# Features Showcase

## A more convenient error handling (inspired by Rust)

```go
import github.com/n-mou/yagul/g

// This
dirInfo := g.Unwrap(os.Stat("."))

// Is equivalent to this
dirInfo, err := os.Stat(".")
if err != nil {
	panic(err)
}

// And this
g.Force(os.Mkdir("tmp_dir", 0755))

// Is equivalent to this
err := os.Mkdir("tmp_dir", 0755)
if err != nil {
	panic(err)
}
```

## A more convenient way of writing Go iterators

```go
import (
	"container/list"
	"fmt"
	"github.com/n-mou/yagul/itertools"
)
package main

// Implementing an iterator for Go's standar library double
// linked list creating a type that implements the 
// itertools.PullIterator interface
type ListIterator struct {
	currentNode *list.Element
}

func newListIterator(l *list.List) *ListIterator {
	return &ListIterator{l.Front()}
}

// Next returns the next element of the iterator and a
// boolean that signals if the iterator has finished.
func (l *ListIterator) Next() (any, bool) {
	if l.currentNode == nil {
		return nil, false
	}
	returnVal := l.currentNode.Value
	l.currentNode = l.currentNode.Next()
	return returnVal, true
}

// Stop is a cleanup function. If the iterator uses resources
// that must be released manually (like closing a file), 
// place tht code here.
func (l *ListIterator) Stop() {}

func main() {
	values := []string{"A", "B", "C", "D", "E"}
	l := list.New()
	for _, i := range values {
		l.PushBack(i)
	}

	// Get a regular iterator (iter.Seq) from any
	// type that implements itertools.PullIterator.
	// Similar functions are implemented for 
	// iterators of key-pair values.
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
```
## Some file management utilities

```go
import "github.com/n-mou/yagul/fs"

// Copy file from source to dest. The file contents
// is loaded in a 32KB buffer so even files that 
// weigh several GBs it won't crush. 
n, err := fs.CopyFile("srcfile", "dstfile")

// Copy srcdir to dstdir and recursively copy all of
// it's files an subdirectories.
err := fs.CopyDir("srcdir", "dstdir")

```

# Why?

I'm relatively new to Go development and I've seen **many devs create a Go-utils repo** containing helper libraries they find useful. Some repos are very extensive and others are a small compilation of simple and selective helpers. However, **none of these Go-utils package is popular enough to became a de-facto standard** (like JQuery was in early Javascript). 

I could make my own compilation but in most of these libraries there are some common functions that are rewritten over and over (like the functional `map()`, `filter()` and `reduce()`). So **instead of reinventing the wheel, I'll just post modules that I missed in those Go-utils repos** and that I think other Go devs might benefit from them.

Thus, this repo is my humble contribution of helper functions and dev patterns that I find useful and not very popular in the Go community instead of a full fledged Go-utils library that aims to be a JQuery for Go. If you like it and use it in your projects, please give a star.

# Roadmap (lack of)

This repo is a compilation of utilities I needed as a Go developer and wasn't able to find in third party projects. There are no projections on when or how will it be expanded or updated. However, **I'll try not to change the type signature of the existing functions** to avoid breaking somebody else's code that might rely on some submodule.

All commited code is documented and tested, and if eventually someone ends using this repo to the point that is willing to contribute, **PRs are appreciated and welcome**.

# License

In case of someone wanting to use some of this modules, this repo is released under the MIT license. So you can copy, alter and reditribute the code at your will. But please, give some credit and leave a link to this software repo somewhere in your own license.