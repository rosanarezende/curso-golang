package main

import "fmt"

func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func calculosMatematicos(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(texto string) string {
		fmt.Println(texto)
		return texto
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	// _, resultadoSubtracao := calculosMatematicos(10, 15)
	// fmt.Println(resultadoSubtracao)

	// resultadoSoma, _ := calculosMatematicos(10, 15)
	// fmt.Println(resultadoSoma)

	resultadoSoma2, resultadoSubtracao2 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2, resultadoSubtracao2)
}
