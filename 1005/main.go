package main

import "fmt"

func main() {

	var num1 float64
	var num2 float64

	fmt.Scan(&num1)
	fmt.Scan(&num2)
	num1 = num1 * 3.5
	num2 = num2 * 7.5

	fmt.Printf("MEDIA = %.5f\n", ((num1 + num2) / 11))
}
