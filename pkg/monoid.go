package gocat

type Num interface {
	int | int64 | uint | uint64 | float64
}

// MonoidSum is monoid for Nums for addition.
type MonoidSum[A Num] struct{ val A }

func unitMonoidSum[A Num]() MonoidSum[A] {
	return MonoidSum[A]{0}
}

func compMonoidSum[A Num](x, y MonoidSum[A]) MonoidSum[A] {
	return MonoidSum[A]{x.val + y.val}
}

// MonoidProd is monoid for Nums for multiplication.
type MonoidProd[A Num] struct{ val A }

func unitMonoidProd[A Num]() MonoidProd[A] {
	return MonoidProd[A]{1}
}

func compMonoidProd[A Num](x, y MonoidProd[A]) MonoidProd[A] {
	return MonoidProd[A]{x.val * y.val}
}

// MonoidArr is monoid for arrays.
type MonoidArr[A any] struct {
	arr []A
}

func unitMonoidArr[A any]() MonoidArr[A] {
	return MonoidArr[A]{[]A{}}
}

func compMonoidArr[A any](x, y MonoidArr[A]) MonoidArr[A] {
	r := []A{}
	r = append(r, x.arr...)
	r = append(r, y.arr...)
	return MonoidArr[A]{r}
}
