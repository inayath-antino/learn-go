package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	// initalized with non zero values
	var intArr [3]int

	fmt.Println(intArr)

	// initalized array during declaration. Initialzing with less number of values than size will fill space with zero value
	var intArr2 = [5]int{1, 2, 3}
	fmt.Println(intArr2)

	// can skip specifying size and leave to compiler to detect
	var intArr3 = [...]int{1, 2}
	fmt.Println("Size of intArr3", len(intArr3))

	// adding values to array is only possible through indexes
	intArr3[0] = 4

}
