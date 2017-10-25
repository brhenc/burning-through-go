package main

import "fmt"

func main() {
	// I refuse to use :=
	//and I refuse what seem to be globally defined iterators 

	//for var j int = 7;  j <= 9; j++ {
	//	fmt.Println(j)
	//}
	//./for-clearer.go:8: var declaration not allowed in for initializer
	//./for-clearer.go:8: syntax error: missing { after for clause
	// why. 

	i := 1 //why doesn't this work? 
	//for i <= 3; i++ {
	//	fmt.Println(i)
	//}
	for i <= 3 {
		fmt.Println(i)
		i++ //there's only i++ it seems
	}
	//for i = 0; i <= 3 {
	//	fmt.Println(i)
	//	i++
	//}
	i = 0
	for i <= 3 {
		fmt.Println(i)
		i++
	}
}
