// Package g holds types and functions that could live as global variables instead
// of a module (currently it implements single line expressions for panicking when an
// error is returned). The g is for "Global".
//
// Currently it only has  functions that serve as syntax sugar for calling functions
// and panicking when an error is found. Although panicking whenever an error is found
// is not a proper way of handling them, sometimes it is the right way to do so, and
// there's no syntax sugar provided from Go. These functions provide a more convenient
// way of handling this case. Like the unwrap() function in Rust's result type.
package g

// Unwrap takes any function that returns a value of any type and an error, panics
// if there's an error and returns the value if there's not. It's a one-line expression
// for the classic "if err != nil {return err}", like the unwrap() function in Rust
// Results. It's a generic function, so theres no need to type cast the value it returns
// and the generic type T can be inferred from the function call, so there's no need to
// specify it. Check the example for it's basic usage.
//
// The function's named Unwrap because it does what unwrap() does in Rust, but there's
// also an Unwrap function in the [errors] package to unwrap errors (this function is
// called unwrap in Rust because it unwraps a result). So maybe Unwrap is not the best
// name but I still haven't found a name that convinces me more (maybe Force, but it's
// already taken). On the other hand, anyone that knows Rust knows the unwrap() function
// so giving this function the same name I can transmit what this function does better
// than any sophisticated doc or example can. So I'm not sure if I should change it's
// name or not.
func Unwrap[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Force takes any function that does not return any value but may return an error and
// panics if an error is returned. It's the one-line equivalent of Unwrap for functions
// that don't return any value. Since the name Unwrap was taken and Go doesn't support
// operator overloading, Force is the best alternative. Check the example for it's basic
// usage.
func Force(err error) {
	if err != nil {
		panic(err)
	}
}
