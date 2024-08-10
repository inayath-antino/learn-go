package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Provide Rating")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Printf("Entered Input - %v", input)

	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rating after added by 1 - ", num+1)
	}
}
