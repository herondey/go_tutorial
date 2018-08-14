package basic3

import "fmt"

//Slicesgo とは
func Slicesgo() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	//var s []int = primes[1:4]
	s := primes[1:4]
	fmt.Println(s)
}
