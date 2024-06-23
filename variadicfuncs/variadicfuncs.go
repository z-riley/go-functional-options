package variadicfuncs

type Thing struct {
	opt1 int
	opt2 string
}

func NewThing(opts ...func(*Thing)) *Thing {
	t := defaultOpts()
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func WithOpt1(n int) func(*Thing) {
	return func(t *Thing) {
		t.opt1 = n
	}
}

func WithOpt2(s string) func(*Thing) {
	return func(t *Thing) {
		t.opt2 = s
	}
}

func defaultOpts() *Thing {
	return &Thing{
		opt1: 0,
		opt2: "default",
	}
}
