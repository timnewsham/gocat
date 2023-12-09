package gocat

// Prod is the product of types A and B.
type Prod[A any, B any] struct {
	fst A
	snd B
}

func prod[A any, B any](x A, y B) Prod[A, B] {
	return Prod[A, B]{x, y}
}

// fst and snd are the projects of Prod[A,B] to A and B.
func fst[A any, B any](p Prod[A, B]) A { return p.fst }
func snd[A any, B any](p Prod[A, B]) B { return p.snd }
