package basic1

import(
	"fmt"
)

var i, j int = 1, 2

func Variableswithinitializers(){
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}