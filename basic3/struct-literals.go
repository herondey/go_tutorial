package basic3

import "fmt"

//Vertex4 とは
type Vertex4 struct {
	X, Y int
}

var (
	v1 = Vertex4{1, 2}
	v2 = Vertex4{X: 1}
	v3 = Vertex4{}
	p  = &Vertex4{1, 2}
)

//Structliterals とは
func Structliterals() {
	fmt.Println(v1, p, v2, v3)
}
