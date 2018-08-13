package basic2

import(
	"fmt"
)

//Defermulti とは
func Defermulti(){
	fmt.Println("counting")

	for i:=0; i<10; i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
}