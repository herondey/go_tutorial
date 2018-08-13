package basic1

import(
	"fmt"
	"math"
)

func Typecomversions(){
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	fmt.Println(x, y, z)
}