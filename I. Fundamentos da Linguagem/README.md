# RESUMO DAS AULAS

## Vantagens de prender Go (Golang)

- **Eficiente**:
  - <ins>compilada</ins>: + rápida do q interpretadas
  - <ins>otimizada para usar + de 1 núcleo do processador</ins>: a linguagem surgiu depois q os processadores já tinham + de 1 núcleo
  - <ins>lida bem com concorrência</ins>
  <br>

- **Simplicidade + robustez**
  - <ins>simples de aprender e entender como JS/Python</ins> (mas elas flexibilidade até demais)
  - <ins>robusta como C#/Java</ins> (mas complexas, até difíceis de aprender), por ser fortemente tipada
  <br>

- **Criada pelo Google**
  - feita pra resolver problemas de uma empresa gigante... e tb mantida por ela.
  <br>

- **Utilizada por grandes empresas**
  - #ex: Google, Facebook, Uber, Twitch, IBM
  - inclusive empresas brasileiras

## Anotações gerais

**Lista de imports**: referencia os pacotes q iremos usar no arquivo
  ```go
  import (
    "fmt"
  )
  ```
<br>

**Terminal**
- pra rodar:
  ```go
  go run <nome do arquivo>
  ```

<br>

## I - Pacote e Módulo

**Pacote**: grupo de arquivos num mesmo diretório, compilados juntos
- uma variável/função/etc declarada em um arquivo será visível para todos os arquivos que tb estiverem nesse pacote
- <ins>main</ins>:  go identifica que esse arquivo é executável
package main
<br>

- apesar de go não ser orientado a objetos (não tendo public, private e protected, por exemplo), nós podemos dizer de determinada função/variável/struct pode ou não ser chamada em outros pacotes
  - letra **Maiúscula**: pode ser acessado de fora do pacote... é pública
    OBS: deve ter uma msg acima dessa função explicando o que ela faz
  - letra **minúscula**: só do próprio pacote

<br>

**Módulo**: Ao lidar com +d1 pacote precisamos criar um módulo (*conj de pcts q contem o projeto -> q vc criou ou externo*)
- a ideia é q o go compile todo esse código do projeto em um lugar só
- é como no package.json do JS (centralizando todas as dependências)
- pra criar o `go.mod`
  ```go
  go mod init <nome do módulo>
  ```

<br>

E pra criar o módulo em si
  ```go
  go build
  ```
- aí ele cria um módulo com o nome que demos anteriormente

<br>

Agora posso rodar o projeto rodando o arquivo diretamente
  ```go
  ./<nome do módulo>
  ```

<ins>OBS</ins>: a medida que for atualizando o arquivo preciso dar o build novamente.
<br>

<ins>OBS2</ins>: o comando `go install` faz o mesmo que o build, mas salva o arquivo main.go na raiz onde foi intalado o go ao invés de ficar no projeto
<br>

<ins>OBS3</ins>: **Instalando pacote EXTERNO**

Na mesma pasta do módulo:

```go
go get <caminho do pacote>
```
- ele adiciona o pacote no `go.mod` e tb cria um arquivo `go.sum`

E pra usar no pacote é só importar

```go
...
import (
	"github.com/badoux/checkmail"
)
...
func main() {
  erro := checkmail.ValidateFormat("devbook@gmail.com")
  fmt.Println(erro)
}
```
<br>

<ins>OBS4</ins>: se eu quiser **LIMPAR** pacotes externo não usados na aplicação

 ```go
go mod tidy
```

<br>

## II - Variáveis

Maneiras de **atribuir** variáveis:

```go
var variavel1 string = "variável 1"

variavel2 := "variável 2"

var (
	variavel3 string = "variável 3"
	variavel4 string = "variável 4"
)

variavel5, variavel6 := "variável 5", "variável 6"
	
const constante1 string = "constante 1"
```

Podemos **reatribuir** variáveis:

```go
variavel5, variavel6 = variavel6, variavel5
```

<br>

## III - Tipos de dados

### Números

<ins>Inteiros</ins>:

```go
	var numero1 int64 = -100000

	var numero2 uint32 = 100000

	// Alias =>>> INT32 = rune
	var numero3 rune = 123456

	// Alias =>>> UINT8 = byte
	var numero4 byte = 123
```

<ins>Reais</ins>:

```go
	var numeroReal1 float32 = 123.45

	var numeroReal2 float64 = 123678.45

	numeroReal3 := 123.45
```

<br>

### Strings

```go
	var texto1 string = "texto"

	texto2 := "texto"

	char := 'A'
```

	// ============

	texto3 := 5

<br>

### Booleanos

```go
	var booleano1 bool
```

<br>

### Erros

```go
var erro error = errors.New("erro interno")
```

<br>

## IV - Funções

```go
func somar(num1 int8, num2 int8) int8 {
	return num1 + num2
}
soma := somar(10, 20)
```


```go
func calculosMatematicos(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}
_, resultadoSubtracao := calculosMatematicos(10, 15)
```


```go
var f = func(texto string) string {
	fmt.Println(texto)
	return texto
}
resultado := f("Texto da função 1")
```

<br>

## V - Operadores

Ver arquivo 
- mesmo que em JS

<br>

## VI - Structs

Criando:

```go
type endereco struct {
	logradouro	string
	numero 			uint8
}

type usuario struct {
	nome 			string
	idade		 	uint8
	endereco	endereco
}
```

<br>

Usando:

```go
enderecoExemplo := endereco{"Rua dos Bobos", 0}


var u usuario
u.nome = "Davi"
u.idade = 21

usuario2 := usuario{"Davi", 21, enderecoExemplo}

usuario3 := usuario{idade: 22}
```

<br>

## VII - ~~Herança~~ (só que não)

Declarando as structs:

```go
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
```

<br>

Usando:

```go
pessoa1 := pessoa{"Rosana", "Rezende", 34, 163}

estudante1 := estudante{pessoa1, "Engenharia", "USP"}
```

<br>

## VIII - Ponteiros

Ponteiro é uma referência de memória

- sem ~~ponteiro~~
  ```go
  var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)
  // 

	variavel2++
	fmt.Println(variavel1, variavel2)
  //
  ```

- com ponteiro
  ```go
  var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
  //

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
  //
  ```

<br>

## IX - Arrays e Slices



<br>

## X - Maps



<br>

## XI - If Else



<br>

## XII - Switch



<br>

## XIII - Loops



<br>

## XIV - Funções Avançadas



<br>

## XV - Métodos



<br>

## XVI - Interfaces



<br>
