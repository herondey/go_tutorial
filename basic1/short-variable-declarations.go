package basic1

import (
	"fmt"
)

func Shortvariabledeclarations(){
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}