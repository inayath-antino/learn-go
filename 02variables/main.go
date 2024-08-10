package main

import "fmt"

// cannot use no var style outside of a function
// someValue := 24

// once change the value
const SECRET = "tokensecret"

func main() {
	fmt.Println("Variables")

	var str string = "hello"
	fmt.Printf("Variable Type - %T and value - %v \n", str, str)

	var isLoggedIn bool = false
	fmt.Printf("Variable Type - %T and value - %v \n", isLoggedIn, isLoggedIn)

	var smallPositiveValue uint8 = 12
	fmt.Printf("Variable Type - %T and value - %v \n", smallPositiveValue, smallPositiveValue)

	// implicit conversion
	var str2 = "hi"
	fmt.Printf("Variable Type - %T and value - %v\n", str2, str2)

	// no var style(walrus operator)
	someValue := 24
	fmt.Printf("Variable Type - %T and value - %v\n", someValue, someValue)

}
