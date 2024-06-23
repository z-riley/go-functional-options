package main

import "fmt"

func main() {
	t1 := NewThing()
	fmt.Printf("t1: %+v\n", t1)

	t2 := NewThing().
		WithOpt1(5).
		WithOpt2("custom")
	fmt.Printf("t1: %+v\n", t2)

}

/// Move to different package

type Thing struct {
	opt1 int
	opt2 string
}

func NewThing() Thing {
	return defaultOpts()
}

func (t Thing) WithOpt1(n int) Thing {
	t.opt1 = n
	return t
}

func (t Thing) WithOpt2(s string) Thing {
	t.opt2 = s
	return t
}

func defaultOpts() Thing {
	return Thing{
		opt1: 0,
		opt2: "default",
	}
}
