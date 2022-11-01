package main

import "fmt"

func main() {
	num1 := 14
	num2 := 24

	sum := num1 + num2
	subtracted := num1 - num2
	multiplied := num1 * num2
	divided := 56 / 8
	modulus := 15 % 2

	fmt.Println(sum, subtracted, multiplied, divided, modulus)

	num1++ // 14 + 1
	num2-- // 24 - 1
	fmt.Println(num1, num2)
}
