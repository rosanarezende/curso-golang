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

<br>

## Anotações gerais

<br>

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

<br>

**Pacote**: grupo de arquivos num mesmo diretório, compilados juntos

- uma variável/função/etc declarada em um arquivo será visível para todos os arquivos que tb estiverem nesse pacote
- <ins>main</ins>: go identifica que esse arquivo é executável
  package main
  <br>

- apesar de go não ser orientado a objetos (não tendo public, private e protected, por exemplo), nós podemos dizer de determinada função/variável/struct pode ou não ser chamada em outros pacotes
  - letra **Maiúscula**: pode ser acessado de fora do pacote... é pública
    OBS: deve ter uma msg acima dessa função explicando o que ela faz
  - letra **minúscula**: só do próprio pacote

<br>

**Módulo**: Ao lidar com +d1 pacote precisamos criar um módulo (_conj de pcts q contem o projeto -> q vc criou ou externo_)

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

<br>

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

<br>

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

## II - Variáveis / Constantes

<br>

Maneiras de **atribuir** <ins>variáveis</ins>:

- só uma:

  ```go
  // tipando
  var variavel1 string = "variável 1"

  // inferência de tipo
  variavel2 := "variável 2"
  ```

<br>

- várias:

  ```go
  var (
  	variavel3 string = "variável 3"
  	variavel4 string = "variável 4"
  )

  variavel5, variavel6 := "variável 5", "variável 6"
  ```

<br>

E <ins>constantes</ins>:

```go
const constante1 string = "constante 1"
```

<br>

<ins>OBS</ins>: Podemos **inverter** o valor das variáveis sem precisar de um auxiliar:

```go
variavel5, variavel6 = variavel6, variavel5
```

- segue as mesmas opções das variáveis

<br>

## III - Tipos de dados

<br>

### Números

<br>

<ins>Inteiros</ins>:

- São 4 tipos, a depender da qtd de bites: int8, int16, int32, int64.
- Ou posso não especificar (**int**), e aí ele usa a arquitetura do seu computador como base.

  ```go
  // ======== INT8 ========
  var numero8 int8 = 123
  // Alias
  var numero8 byte = 123
  	// OBS: se eu coloco um número muito grande nele dá erro
  	//constant 1230000000000000 overflows byte


  // ======== INT16 ========
  var numero16 int16 = 100


  // ======== INT32 ========
  var numero32b int32 = 123456
  // Alias
  var numero32b rune = 123456


  // ======== INT64 ========
  var numero64 int64 = -1000000000

  ```

<br>

Podemos atribuir números negativos, como vimos, mas tb podemos colocar um **int sem sinal** (unsygned int)

    ```go
    var numero32a uint32 = 123456
    ```

- não aceita atribuição de número negativo

<br>

<ins>Reais</ins>:

- só tem dois tipos: float32 e float64

```go
	var numeroReal1 float32 = 123.45

	var numeroReal2 float64 = 123678.45

	// ele vai inferir pelo computador... diz float, mas só ele pode usar assim
	numeroReal3 := 123.45
```

<br>

### Strings

```go
	var texto1 string = "texto"

	texto2 := "texto"
```

<br>

<ins>OBS</ins>: Go não tem ~~char~~.

- um caractére só com aspas simples ele converte pra número

  ```go
  	char := 'A'
  	// converte para o número da tabela asp... nesse caso 65
  	// se colocar +d1 caracter dá erro
  ```

<br>

<ins>OBS2</ins>:

```go
texto3 := 5
```

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

### Array

<ins>Criando</ins>

```go
var array1 [5]string
array1[0] = "Posição 1"

array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

array3 := [...]int{1, 2, 3, 4, 5}
```

### Slice

<ins>Criando</ins>

- do zero:

```go
slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
```

- de um array

  ```go
  slice2 := array2[1:3]
  ```

  - e se eu alterar o array?
    ```go
    array2[1] = "Posição Alterada"
    fmt.Println(slice2)
    //
    ```

<ins>Adicionando</ins>

```go
	slice = append(slice, 18)
```

### Arrays internos

<ins>Criando</ins>

```go
slice3 := make([]float32, 10, 11)
fmt.Println(slice3)
//
```

<ins>Adicionando</ins>

```go
slice3 = append(slice3, 5)
slice3 = append(slice3, 6)
```

- fica de olho:
  ```go
  fmt.Println(slice3)
  fmt.Println(len(slice3)) // length
  fmt.Println(cap(slice3)) // capacidade
  ```

<br>

## X - Maps

<ins>Criando</ins>

    ```go
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
    ```

<br>

<ins>Deletando</ins>

```go
	delete(usuario2, "nome")
```

<br>

<ins>Adicionando</ins>

```go
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}
```

<br>

## XI - If Else

```go
	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}


	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("Numero menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
```

<br>

## XII - Switch

```go
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Domingo"
	default:
		return "Número Inválido"
	}
}
dia := diaDaSemana(200)
fmt.Println(dia)
//


func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Domingo"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana
}
dia2 := diaDaSemana2(1)
fmt.Println(dia2)
//
```

<br>

## XIII - Loops

```go
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)
```

```go
	for j := 0; j < 10; j += 5 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
```

<br>

<ins>Range</ins>

```go
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
```

<br>

<ins>Infinito</ins>

```go
	// for {
	// 	fmt.Println("Executando infinitamente")
	// 	time.Sleep(time.Second)
	// }
```

<br>

## XIV - Métodos

```go
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
```

<br>

## XV - Interfaces

### FORMAS

```go

```

<br>

### TIPO GENÉRICO

```go
	func generica(interf interface{}) {
		fmt.Println(interf)
	}

	generica("String")
	//

	generica(1)
	//

	generica(true)
	//
```

<br>

```go
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)
```

<br>
