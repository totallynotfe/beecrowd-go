package main

import "fmt"

func main() {
	var num1, num2, num3 float64
	fmt.Scan(&num1, &num2, &num3)
	num1 *= 2
	num2 *= 3
	num3 *= 5
	fmt.Printf("MEDIA = %.1f\n", ((num1 + num2 + num3) / 10))
}
