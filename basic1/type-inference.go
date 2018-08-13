package basic1

import (
	"fmt"
	"math/cmplx"
)

func Typeinference(){
	//v := 42
	//v := 24.3
	v := cmplx.Sqrt(100)
	fmt.Printf("v is of type %T\n", v)
}