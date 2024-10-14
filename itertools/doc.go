/*
Package itertools provides an alternative way of writing Go iterators in a more
convenient way and some other helper functions to handle iterators.

# Why?

When Go released the 1.23 version with the iterators I was getting started with
the language and I didn't know about all the controversy. After looking the type
signature and the [iter] package I learned to create them but found the syntax
very unconvenient, and I thought that if the pattern was as a pull iterator
instead of a push iterator (and inspired of how it's done in Python and Java) it
would be easier to create them. And thus, this library was created.

# How Do I Make Iterators With this Library?

There are 2 new interfaces ([PullIterator] and [PullIterator2]) that define
2 functions (Next and Stop). Any type that implements those 2 functions
is a valid [PullIterator] or [PullIterator2] and can betransformed into an
[iter.Seq] or [iter.Seq2] types by the [PullToPush] or [PullToPush2] function,
which then, can be used as an argument in the "for i := range" syntax block.

	// Type signatures of Next
	// In PullIterator, Next returns a value and a boolean and it will be transformed
	// into a iter.Seq[V] type
	func Next() (V, bool)
	// In PullIterator2, Next returns a key-value pair and a boolean and it will
	// be transformed into a iter.Seq2[K, V] type
	func Next() (K, V, bool)

The Next function returns the next value of the collection (or the next key-value
pair in the case of [PullIterator2]) and a boolean each time is called. The boolean
states if the returned value (or key-value pair) is valid. When returning an element
o the collection, it will always be true. Once the last element is returned, future
function calls no matter what value or key-value pair return as long as the bool is
false (signaling that the iterator has finished and future calls to Next() will
return invalid results). This can be achieved by implementing a struct with a field
that holds the current element of the collection and updates it on each Next() call.
Notice that since each call mutates the struct, so the Next function must have a
pointer receiver instead of a value receiver.

	// Type signture of Stop
	func Stop()

The Stop function is a cleanup function, it's called when the iterator's finished (when
Next() starts returning false) and if the program breaks from the for loop before the
iterator is exhausted. Usually it's an empty function but in case the iterator relies
on resources that need to be closed (such as files or network or database connections),
that code must be in this function.

Finally, when calling [PullToPush] or [PullToPush2] you must pass a pointer of the
struct implementing those methods instead of the struct itself. Since Next and Close
methods have likely a pointer receiver instead of a value receiver, the Go compiler
can't infer that the struct implements the [PullIterator] or [PullIterator2] interface:

	// Asuming the struct implements Next() and Stop() with pointer receivers, this
	// will not compile
	iterator := itertools.PullToPush(iterStruct{})
	// but this will
	iterator := itertools.PullToPush(&iterStruct{})

Check the examples to see this library in action.

# What's a Pull Based Iterator?

Pull and push based are theoretical concepts that describe how iterators behave
and are defined. For this particular case, it means that instead of writing a
function that traverses all the collection calling yield() (a callback
that receives the current element and runs whatever is inside the for {} block)
you write a function that returns the next element of the collection every time
it's called and a boolean signaling if the iterator has finished or not.

This style of iterators are more convenient for some particular cases, that's
why Go implements the [iter.Pull] and [iter.Pull2] function, that transforms any
iterator ([iter.Seq] or [iter.Seq2]) into a pull based iterator returning
the functions next() and stop(). This library uses the same concept with the
same type signature to go in the opposite direction (create an [iter.Seq] or
[iter.Seq2] from a pull based iterator).
*/
package itertools
