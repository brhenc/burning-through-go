// where is my sizeof()?
package main

import "fmt"

func main() {

	var str string = "initial" //can't use "string" as variable.
	fmt.Println(str)
	var b, c int = 1, 2 //var int b, c would be so much clearer. This is insane
	//var not_retarded int = 1
	//var still_not_retarded int = 2 //I'll stick to one declaration per line. thx.
	fmt.Println(b, c)


	var d = true //what happened to bool? 
	//var ded bool = true //ok here it is. Is it one byte or one bit? (probably neither)
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"	//I hate pascal and I hate := . this is bad, you're redeclaring :=
			//which is worse than pascal and := . I'll stick to var thanks
			//because I don't want to confuse people, it's var in javascript
	fmt.Println(f)
}
//https://stackoverflow.com/questions/21743841/how-to-avoid-annoying-error-declared-and-not-used-from-golang
//go 
//https://golang.org/doc/faq#unused_variables_and_imports
//What moron needs more than a warning about unused variables? It should be a warning.
//Thank you for forcing me to use training wheels.
