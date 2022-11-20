package main

import "fmt"

func main() {
	const PI = 3.14159
	var r float64
	fmt.Scanf("%f", &r)
	fmt.Printf("A=%.4f\n", (PI * (r * r)))
}
