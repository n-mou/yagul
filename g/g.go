// Package g holds types and functions that could live as global variables instead
// of a module. The g is for "Global". Currently it only has functions that serve
// as syntax sugar for calling and panicking when an error is found in a single line.
//
// # Warning
//
// This functions are meant to ease testing and prototyping but panicking hardly ever
// is the way to go when handling errors. If you're developing a third party library no
// function should ever panic when an error is found. It's preferable to return the
// error and let the program using that library handle it. And if you're developing an
// application, it's advised to know what types of errors can the libraries and APIs
// used return and handle each case leaving the panic only when an unrecoverable error
// is expected and even then, there's some controversy about it.
package g

// Must takes any function that returns a value of any type and an error, panics
// if there's an error and returns the value if there's not. It's a one-line expression
// for the classic "if err != nil {return err}", like the unwrap() function in Rust
// Results. It's a generic function, so theres no need to type cast the value it returns
// and the generic type T can be inferred from the function call, so there's no need to
// specify it. Check the example for it's basic usage.
func Must[T any](v T, err error) T {
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
