package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true // infers type
	fmt.Println(d)

	var e int // will be set to 0 initialised value
	fmt.Println(e)

	// var f string = "apple"
	f := "apple" // shorthand declare & initialise
	fmt.Println(f)
}
