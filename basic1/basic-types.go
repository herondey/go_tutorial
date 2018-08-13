package basic1

import (
	"fmt"
	"math/cmplx"
)

var tobe bool
var maxint uint64 = 1<<64 - 1
var z complex128 = cmplx.Sqrt(-5 + 12i)

func Basictypes() {
	fmt.Printf("Type: %T value: %v\n", tobe, tobe)
	fmt.Printf("Type: %T Value: %v\n", maxint, maxint)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
