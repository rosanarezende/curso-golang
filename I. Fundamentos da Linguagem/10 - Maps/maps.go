package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":				"Rosana",
		"sobrenome":	"Rezende",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":	"Pedro",
			"último":		"Lucas",
		},
		"curso": {
			"nome":		"Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
	fmt.Println(usuario2)
}