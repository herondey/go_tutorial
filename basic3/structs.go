package basic3

import (
	"fmt"
)

//Vertex とは
type Vertex struct {
	X int
	Y int
}

//Structs とは
func Structs() {
	fmt.Println(Vertex{1, 2})
}
