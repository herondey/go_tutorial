package basic2

import (
	"fmt"
	"math"
)

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		//return v
		lim = v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

//Ifandelse とは
func Ifandelse() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}
