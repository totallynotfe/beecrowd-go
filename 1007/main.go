package main

import "fmt"

func main() {
	var num1, num2, num3, num4 int
	fmt.Scan(&num1, &num2, &num3, &num4)
	num1 *= num2
	num3 *= num4
	fmt.Printf("DIFERENCA = %d\n", (num1 - num3))
}
