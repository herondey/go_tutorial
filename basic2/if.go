package basic2

import(
	"fmt"
	"math"
)

func sqrt(x float64) string{
	if x < 0{
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//Ifgo とは
func Ifgo(){
	fmt.Println(sqrt(2), sqrt(-4))
}