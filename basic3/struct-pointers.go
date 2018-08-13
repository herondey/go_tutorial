package basic3

import "fmt"

//Vertex3 とは
type Vertex3 struct{
	X int
	Y int
}

//Structpointers とは
func Structpointers(){
	v := Vertex3{1, 2}
	p := &v
	(*p).X = 1e9
	fmt.Println(v)
}