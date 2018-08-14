package basic3

import (
	"fmt"
)

//Appendgo とは
func Appendgo() {
	var s []int
	printSlice3(s)

	s = append(s, 0)
	printSlice3(s)

	s = append(s, 1)
	printSlice3(s)

	s = append(s, 2, 3, 4)
	printSlice3(s)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
