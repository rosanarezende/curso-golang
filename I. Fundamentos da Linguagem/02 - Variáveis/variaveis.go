package main

import "fmt"

func main() {
	var variavel1 string = "variável 1"
	variavel2 := "variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variável 3"
		variavel4 string = "variável 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variável 5", "variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
