package main

import "fmt"

func main() {
	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("Numero menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
