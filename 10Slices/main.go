package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices")

	// slices under the hood are arrays

	// uninitalized slice
	var intSlice []int
	fmt.Println("Uninitailized slice", intSlice, "Capacity", cap(intSlice), "Length", len(intSlice))

	// initializing slice with empty brackets
	var intSlice1 = []int{}
	fmt.Println("Initialized Slice with empty brackets", intSlice1, "capacity", cap(intSlice1), "length", len(intSlice1))

	var intSlice2 = []int{1, 2, 3}
	fmt.Println("Initialized slice", intSlice2, "capacity", cap(intSlice2), "length", len(intSlice2))

	// adding values to the end of the slice
	intSlice2 = append(intSlice2, 8)

	// adding multiple values to the end of the slice
	intSlice2 = append(intSlice2, 54, 23)
	fmt.Println(intSlice2)

	// slicing
	intSlice2 = intSlice2[3:]
	fmt.Println(intSlice2)

	// declaring using make
	intSlice3 := make([]int, 3, 5)

	// since size is specified here we can access it through indexes only upto the size, accessing more than the size will result in out of bound error
	intSlice3[0] = 780
	intSlice3[1] = 1234
	// intSlice3[2] = 546

	fmt.Printf("Memory Address of intSlice3 %p length %v capacity %v\n", &intSlice3, len(intSlice3), cap(intSlice3))

	// this will throw index out of bound error
	// intSlice3[3] = 345

	// asc sorting
	// fmt.Println("Before sorting intSlice3", intSlice3)
	// sort.Ints(intSlice3)
	// fmt.Println("After ASC", intSlice3)

	// reverse sorting
	// fmt.Println("Before sorting intSlice3", intSlice3)
	// sort.Sort(sort.Reverse(sort.IntSlice(intSlice3)))
	// fmt.Println("After DESC", intSlice3)
}
