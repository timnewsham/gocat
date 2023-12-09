package gocat

// Morph is a morphism from A to B.
type Morph[A any, B any] interface {
	Apply(A) B
}

type Cat interface {
	// `interface method must have no type parameters`
	//id[A any]() Morph[A, A]
	//comp[A any, B any, C any](g Morph[B, C], f Morph[A, B]) Morph[A, C]
}
