package basic2

import (
	"fmt"
)

//Forgo とは
func Forgo() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
