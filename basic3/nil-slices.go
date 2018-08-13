package basic3

import "fmt"

//Nilslices とは
func Nilslices(){
	var s[]int
	fmt.Println(s, len(s), cap(s))
	if s == nil{
		fmt.Println("nil")
	}
}