package basic1

import(
	"fmt"
)

func swap(x, y string)(string, string){
	return y, x
}

func Multipleresults(){
	a, b := swap("Nishizawa", "Takamasa")
	fmt.Println(a, b)
}