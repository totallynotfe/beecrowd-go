package main

import "fmt"

func main() {
	var id, hours int
	var mph float64
	fmt.Scan(&id, &hours, &mph)
	res := mph * float64(hours)
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", id, res)
	return
}
