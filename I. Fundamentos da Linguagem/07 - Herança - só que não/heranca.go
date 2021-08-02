package main

import "fmt"

type pessoa struct {
	nome			string
	sobrenome	string
	idade 		uint8
	altura 		uint8
}

type estudante struct {
	pessoa
	curso 		string
	faculdade string
}

func main(){
	pessoa1 := pessoa{"Rosana", "Rezende", 34, 163}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Engenharia", "USP"}
	fmt.Println(estudante1)
}
