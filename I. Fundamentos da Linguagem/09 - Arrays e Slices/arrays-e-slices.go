package main

import "fmt"

func main() {

	// ======== ARRAY ========

	// Criando
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)


	// ======== SLICE ========

	// Criando
	
	// - do zero
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// - de um array
	slice2 := array2[1:3]
	fmt.Println(slice2)

		//      - e se eu alterar o array?
		array2[1] = "Posição Alterada"
		fmt.Println(slice2)


	// Adicionando
	slice = append(slice, 18)
	fmt.Println(slice)


	// ======== ARRAYS INTERNOS ========

	// Criando
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	// Adicionando
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
}
