package main

import "fmt"

func main() {
	var id, qtd int;
	var unitary_value, result float64;
	result = 0;
	for i := 0; i < 2; i++ {
		fmt.Scanf("%d%d%f", &id, &qtd, &unitary_value)
		result += float64(qtd) * unitary_value
	}
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", result)
}