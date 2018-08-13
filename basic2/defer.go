package basic2

import(
	"fmt"
)

//Defergo とは
func Defergo(){
	defer fmt.Println("world")

	fmt.Println("hello")
}