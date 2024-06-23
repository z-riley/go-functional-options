package main

import (
	"fmt"

	"github.com/zac460/go-functional-options/methodchain"
	"github.com/zac460/go-functional-options/variadicfuncs"
)

func main() {

	// Method chaining

	t1 := methodchain.NewThing()
	fmt.Printf("Default options: %+v\n", t1)

	t2 := methodchain.NewThing().
		WithOpt1(5).
		WithOpt2("custom")
	fmt.Printf("Custom options: %+v\n", t2)

	// Variadic functions

	t3 := variadicfuncs.NewThing()
	fmt.Printf("Default options: %+v\n", t3)

	t4 := variadicfuncs.NewThing(
		variadicfuncs.WithOpt1(10),
		variadicfuncs.WithOpt2("custom"),
	)
	fmt.Printf("Custom options: %+v\n", t4)

}
