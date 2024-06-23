package main

import (
	"fmt"

	"github.com/zac460/go-functional-options/methodchain"
)

func main() {
	t1 := methodchain.NewThing()
	fmt.Printf("Default options: %+v\n", t1)

	t2 := methodchain.NewThing().
		WithOpt1(5).
		WithOpt2("custom")
	fmt.Printf("Custom options: %+v\n", t2)
}
