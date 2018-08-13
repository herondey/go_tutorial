package basic2

import (
	"fmt"
)

//Forcontinued とは
func Forcontinued() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
