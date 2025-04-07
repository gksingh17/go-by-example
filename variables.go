package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	// can declare multiple values at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// no initialization will result with 0 value
	var e int
	fmt.Println(e)

	// short declaration, sytanx only available inside function
	f := "gunjan"
	fmt.Println(f)
}
