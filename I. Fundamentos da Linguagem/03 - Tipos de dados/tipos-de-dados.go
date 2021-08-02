package main

import (
	"errors"
	"fmt"
)

func main() {
	// ============ Números inteiros ===============

	var numero1 int64 = -100000
	fmt.Println(numero1)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	// == Alias ==
	// INT32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// == Alias ==
	// UINT8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)


	//  ============ Números reais ============
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123678.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)


	// ============ Strings ============

	var texto1 string = "texto"
	fmt.Println(texto1)

	texto2 := "texto"
	fmt.Println(texto2)

	char := 'A'
	fmt.Println(char)

	// ============

	texto3 := 5
	fmt.Println(texto3)


	// ============ Booleanos ============
	var booleano1 bool
	fmt.Println(booleano1)

	
	// ============ Erro ============
	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}
