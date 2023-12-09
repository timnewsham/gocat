package gocat

// Func[A,B] is a golang function from A to B.
type Func[A any, B any] struct {
	F func(A) B
}

// Func[A, B] is a morphism.
func (f Func[A, B]) Apply(x A) B { return f.F(x) }

// id is the identity Func.
func idFunc[A any]() Func[A, A] {
	return Func[A, A]{func(x A) A { return x }}
}

// comp is composition of two Funcs, g after f.
func compFunc[A any, B any, C any](g Func[B, C], f Func[A, B]) Func[A, C] {
	return Func[A, C]{
		func(x A) C { return g.Apply(f.Apply(x)) },
	}
}
