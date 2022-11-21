package main

import (
	"fmt"
	"math"
)

func main() {
	var pi float64 = 3.14159
	var radius float64
	fmt.Scan(&radius)
	fmt.Printf("VOLUME = %.3f\n", (4.0/3 * pi * math.Pow(radius, 3)))
}