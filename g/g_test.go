package g_test

import (
	"fmt"
	"os"

	"github.com/n-mou/yagul/g"
)

func ExampleMust() {
	// This is equivalent to:
	// val, err := os.Stat(".")
	// if err != nil {
	//     panic(err)
	// }
	val := g.Must(os.Stat("."))

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
