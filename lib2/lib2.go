package lib2

func Foo() {}

type Circle struct{}

func (Circle) Foo() {}

type Seat struct {
	Description string

	Legs int

	Foo string
}

type Table struct {
	Description string

	Seats int

	Foo string

	Extra struct {
		Foo string
	}
}
