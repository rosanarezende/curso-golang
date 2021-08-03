package main

import (
	"fmt"
	"time"
)

func main(){
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)


	for j := 0; j < 10; j += 5 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	// =========== RANGE ===========
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}


	nomes := [3]string{"João", "Davi", "Lucas"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}


	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// =========== INFINITO ===========
	// for {
	// 	fmt.Println("Executando infinitamente")
	// 	time.Sleep(time.Second)
	// }
}
