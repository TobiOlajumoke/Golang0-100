package main

import (
	"fmt"
	"time"
)

func main() {
	var message string = "Go is amazing"
	fmt.Println(message)

	// Whole numbers - eg 14
	var number int = 14
	fmt.Println(number)

	// real number - 1.43
	var realNum float32 = 1.43
	fmt.Println(realNum)

	// characters - s, !
	var ch rune = 's'
	fmt.Println(ch)

	// boolean - true or false
	var isTrue bool = true
	fmt.Println(isTrue)

	var result time.Time = time.Now()
	fmt.Println(result)

}
