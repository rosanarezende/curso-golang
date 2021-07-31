// pacote -> grupo de arquivos num mesmo diretório, compilados juntos
// uma variável/função/etc declarada em um arquivo será visível para todos os arquivos que tb estiverem nesse pacote
// main ->  go identifica que esse arquivo é executável
package main

// lista de imports -> referencia os pacotes q iremos usar no arquivo
import (
	"fmt"
)

func main() {
	fmt.Println("Olá mundo!")
}

// pra rodar o arquivo, no terminal: 
	// go run <nome do arquivo>
