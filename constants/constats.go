package main

import "fmt"
import "math" 

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n //begs the question, how do I do x^y ?
	fmt.Println(d) // honestly? are you a math major so this is what the res is?

	fmt.Println(int64(d))	//wait, still, begs the question, what kind of 
				//precision do constats have? 

	fmt.Println(math.Sin(n)) // is it Rads, degrees, or grads? 

	fmt.Println("Actually clarify what's going on here")

	const ne = 500
	const de = 5e2 //Honestly, I've never used e. Real men use scientific notation,
			//but don't force it on mortals 
	fmt.Println("I expect 5 here: ", ne/de)

	fmt.Println("I expect degrees, let's do sin(90) = 1 :", math.Sin(1))
	//It's radians of course. ... I like Python more. has both.
	fmt.Println("It's radians... sin(3.14159) = 0 ) :", math.Sin(3.1459))

}
