/*
Package list reimplements the double linked list of Go's standard library using
generics. It's API compatible with the original list.List and list.Element and
it also has some extra quality of life methods (eg: 2 iterators). Using a
generic type T instead of the type any avoids unconvinient type casts every time
an element from the list is retrieved.

# Why?

I needed a linked list for another project and the implementation in Go
standard library holds values of type any. This means any time I'm traversing
the list I have to type cast each node when retrieved. This was the only way
of implementing a list before addition of generics. I find far more convenient
if this data structures held values of a generic type instead of type any but
I guess this part of the standard library is not relevant enough to receive
an update. So, since I needed it I did it myself.

# API Compatible

Both List[T] and Element[T] implement the same funcs with an almost identical
type signature (replacing type any with the generic T). Thus, 99.9% of the Go
code that uses container/list types is reusable. The only exception is the
list.New() func. It creates and returns a new empty list, with this
implementation it's required to declare which type will T adopt.

	// Standard library implementation
	l := list.New()
	// Own implementation
	l := list.New[string]()

Since it's impossible to reimplement the list in a way where this function call
didn't need to be changed I took the opportunity to add the possibility of
initializing a list with some elements. The original New() func takes no args
but this implentation takes a variadic number of args of type T:

	// Initialize an empty list.
	l := list.New[string]()

	// Initialize a list with some elements, notice that in this case it's not necessary to specify
	// the string type since it's inferred from the args.
	l := list.New("A", "B", "C")

# Extra methods

Apart from the methods of the standard library version, these other ones were created to add some
quality of life improvements:

	// Regular iterator:
	for i := range list.Iterator() {
		fmt.Println(i) // i is of type T
	}

	// Backwards iterator, useful to traverse a regular list as a FILO data structure:
	for i := range list.BackwardsIterator() {
		fmt.Println(i) // i is of type T
	}

	// Get a slice of type T holding a copy af all list elements:
	s := list.ToSlice()

	// Finally, both List and Element types have a string representation:
	listRep := list.String()
	elemRep := elem.String()
*/
package list
