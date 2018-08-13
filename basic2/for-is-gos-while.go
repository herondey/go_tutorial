package basic2

import (
	"fmt"
)

//Forisgoswhile とは
func Forisgoswhile(){
	sum := 1
	for sum<1000 {
		sum += sum
	}
	fmt.Println(sum)
}