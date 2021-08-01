package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")

	auxiliar.Escrever()

	// chamar o escrever2 daria erro, pois está com letra minúscula
	// auxiliar.escrever2()

	// chamando função de pacote externo
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
