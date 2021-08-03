package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("Salvando os dados de %s no BD\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Usuario 2", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println("É maior de idade? ", maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println("Nova idade: ", usuario2.idade)
}
