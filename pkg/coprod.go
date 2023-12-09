package gocat

// CoProd is the coproduct of types A and B.
// Invariant: exactly one of `left` or `right` is not nil.
type CoProd[A any, B any] struct {
	left  *A
	right *B
}

// injL injects A into CoProduct[A, B].
func injL[A any, B any](x A) CoProd[A, B] {
	return CoProd[A, B]{&x, nil}
}

// injR injects B into CoProduct[A, B].
func injR[A any, B any](y B) CoProd[A, B] {
	return CoProd[A, B]{nil, &y}
}

// destroy applies destrL or destrR to CoProduct[A, B] to get C.
func destroy[A any, B any, C any](cp CoProd[A, B], destrL Morph[A, C], destrR Morph[B, C]) C {
	if cp.left != nil {
		return destrL.Apply(*cp.left)
	} else {
		return destrR.Apply(*cp.right)
	}
}
