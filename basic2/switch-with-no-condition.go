package basic2

import (
	"fmt"
	"time"
)

//Switchwithnocondition とは
func Switchwithnocondition(){
	t :=time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Goog evening")
	}
}