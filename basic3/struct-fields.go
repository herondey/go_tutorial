package basic3

import(
	"fmt"
)

//Vertex2 とは
type Vertex2 struct{
	X int
	Y int
}

//Structfields とは
func Structfields(){
	v := Vertex2{1, 2}
	v.X = 4
	fmt.Println(v.X)
}