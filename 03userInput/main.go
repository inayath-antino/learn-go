package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter rating")

	// comma ok or comma err
	// ReadString takes in '\n' at the end as well
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error occured %v", err)
	} else {
		fmt.Println("Thanks for input", input)
	}

}
