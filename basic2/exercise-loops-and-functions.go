package basic2

import(
	"fmt"
)

//Sqrt とは
func Sqrt(x float64) float64{
	var z = 1.0
	
	for i:=0; i<10; i++ {
		if x - z*z > 0{
			return z
		}
	}


	return z
}

//Exerciseloopsandfunctions とは
func Exerciseloopsandfunctions(){
	fmt.Println(Sqrt(2))
}
