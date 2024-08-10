package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("05 Handling TIme")

	currentTime := time.Now()

	fmt.Println(currentTime)

	fmt.Println(currentTime.Clock())

	// params - year, month, days
	fmt.Println(currentTime.AddDate(-1, 2, 3))

	// for formatting, string to use is a fixed date(2nd January 2006 Monday, 15h 04min 05sec)
	fmt.Println(currentTime.Format("Mon 06-01-02 15h 04m 05s"))

	// to get any location
	fmt.Println(time.LoadLocation("America/New_York"))

	// creating date
	fmt.Println(time.Date(2023, time.August, 8, 0, 0, 0, 0, time.Local))
}
