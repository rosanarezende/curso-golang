// não necessariamente precisa ser o mesmo nome da pasta
package auxiliar

import "fmt"

// com letra maiuscula ele pode ser acessado de fora do pacote
// Funções exportáveis devem conter uma msg explicativa
	// ex: Escrever registra uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
