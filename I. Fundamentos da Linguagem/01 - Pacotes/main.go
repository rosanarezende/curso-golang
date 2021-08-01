package main

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	auxiliar.Escrever()

	// chamar o escrever2 daria erro, pois está com letra minúscula
	// auxiliar.escrever2()
}
