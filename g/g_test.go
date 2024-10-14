package g_test

import (
	"fmt"
	"os"
	"yagul/g"
)

func ExampleUnwrap() {
	// This is equivalent to:
	// val, err := os.Stat(".")
	// if err != nil {
	//     panic(err)
	// }
	val := g.Unwrap(os.Stat("."))

	fmt.Println(val.IsDir())
	// Output: true
}

func ExampleForce() {
	// This is equivalent to:
	// err = os.Mkdir("tempdir", 0755)
	// if err != nil {
	//     panic(err)
	// }
	g.Force(os.Mkdir("tempdir", 0755))
}
