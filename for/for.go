package main

import "fmt"

func main() {

	var i = 1 //... so, I can only use i once? ... 
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { // ++j and j++ in go. also, := again, required for C for
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue //not a fan of contrived examples. 
		}
		fmt.Println(n)
	}
}
