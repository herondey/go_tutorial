package basic1

import "fmt"

//Pi は円周率ね
const Pi = 3.14

func Constants(){
	const World = "世界"
	fmt.Println("hello ", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}