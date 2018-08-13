package basic1

import(
	"fmt"
)

func split(sum int) (x, y, z int){
	x = sum * 4 / 9
	y = sum - x
	z = sum
	return
}

func Namedresult(){
	fmt.Println(split(9))
}