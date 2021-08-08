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
	#ex: se o PC é 64bits, cria um int64

  ```go
  // ======== INT8 ========
  var numero8 int8 = 123
  // Alias
  var numero8 byte = 123
  	// OBS: se eu coloco um número muito grande nele dá erro
  	//constant 1230000000000000 overflows byte


  // ======== INT16 ========
  var numero16 int16 = 1000


  // ======== INT32 ========
  var numero32b int32 = 123456
  // Alias
  var numero32b rune = 123456


  // ======== INT64 ========
  var numero64 int64 = -1000000000

  ```

<br>

Podemos atribuir números negativos, como vimos, mas tb podemos colocar um **int sem sinal** (unsygned int): UINT

  ```go
  var numero32a uint32 = 123456
  ```

- não aceita atribuição de ~~número negativo~~

<br>

<ins>Reais</ins>:

- só tem dois tipos: float32 e float64

```go
var numeroReal1 float32 = 123.45

var numeroReal2 float64 = 123678.45
```

<br>

Não podemos declarar só ~~float~~, mas se não declaro nada ele infere

```go
// ele vai inferir pelo computador... diz float, mas só ele pode usar assim
numeroReal3 := 123.45
```

<br>

<ins>OBS</ins>: **valor inicial/0**: 0
- é o valor atribuído a uma variável <ins>não inicializada</ins>

```go
var numero int16
// é 0
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
  // converte para o número da tabela ask... nesse caso 65
  // se colocar + de 1 caracter dá erro
  ```

<br>

<ins>OBS2</ins>: **valor inicial/0**: string vazia

```go
var texto3 string
// se eu rodo ela agora é uma: string em branco
```

<br>

### Booleanos

```go
var booleano1 bool = true
```

<br>

<ins>OBS</ins>: **valor inicial/0**: false

<br>

### Erros

- é necessário usar uma lib interna/nativa do go: errors

```go
var erro error = errors.New("erro interno")
```

<br>

<ins>OBS</ins>: **valor inicial/0**: `<nil>`

```go
var erro2 error
// nil
```

<br>

## IV - Funções

- tb é um "tipo" em go

<br>

Forma mais comum

```go
func somar(num1, num2 int8) int8 {
	return num1 + num2
}

// no main():
soma := somar(10, 20)
// 30
```
- quando os tipos dos parâmetros são iguais posso declarar só o último

<br>

Posso igualar uma variável a uma função

```go
var f = func(texto string) string {
	fmt.Println(texto)
	return texto
}
resultado := f("Texto da função 1")
```
- o tipo de **f** fica `func(texto string) string`
<br>
- poderia ser uma função sem retorno, e **f** teria como tipo `func(texto string)`

	```go
	var f = func(texto string) {
		fmt.Println(texto)
	}
	```

<br>

A função pode **ter + de 1 retorno**

```go
func calculosMatematicos(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}

resultadoSoma2, resultadoSubtracao2 := calculosMatematicos(20, 15)
// 25 -5

resultadoSoma, _ := calculosMatematicos(10, 15)
// 25

_, resultadoSubtracao := calculosMatematicos(10, 15)
// -5
```
- isso é muito usado quando lidamos com erro

<br>

## V - Operadores

São os mesmo que em JS

<br>

### ARITMÉTICOS 

```go
soma := 1 + 2
subtracao := 1 - 2
divisao := 10 / 4
multiplicacao := 10 * 5
restoDaDivisao := 10 % 2
```

<br>

<ins>OBS</ins>: Não podemos fazer ~~operações com tipos diferentes~~, mesmo sendo números

```go
var numero1 int16 = 10
var numero2 int32 = 25
soma := numero1 + numero2
// dá erro
```

<br>

### ATRIBUIÇÃO 

```go
var variavel1 string = "String"
variavel2 := "String2"
```

<br>

###  OPERADORES RELACIONAIS

```go
1 >= 2 	// false
1 > 2 	// false
1 == 2 	// false
1 <= 2 	// true
1 > 2 	// false
1 < 2 	// true
1 != 2 	// true
```

<br>

### OPERADORES LÓGICOS

```go
verdadeiro, falso := true, false
verdadeiro && falso 	// false (todas verdadeiras?)
verdadeiro || falso 	// true (uma verdadeira?)
!verdadeiro 		// false
!falso 			// true
```

<br>

###  OPERADORES UNÁRIOS
- só interage com 1 variável por vez

```go
	numero := 10
	
	numero++ // 11
	numero-- // 10
	
	numero += 15 // (numero = numero + 15) => 25
	numero -= 20 // (numero = numero - 20) => 5
	numero *= 3 // (numero = numero * 3) => 15
	numero %= 10 // (numero = numero % 10) => 5
	numero /= 2  // (numero = numero / 2) => 2
```

<br>

### ~~Ternário~~

<br>

Não existe ternário em go
- tem q usar if else

	```go
	numero := 10
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	// numero => Maior que 5
	```

<br>

## VI - Structs

É uma **coleção de campos**, cada qual com *nome* e *tipo*.

<br>

<ins>Criando</ins>:

```go
type endereco struct {
	logradouro	string
	numero		uint8
}

type usuario struct {
	nome		string
	idade		uint8
	endereco	endereco
}
```

<br>

<ins>OBS</ins>: **valor inicial/0**: joga o valor 0 *de cada um dos valores*

```go
var end endereco
// { 0}

var u usuario
// { 0 { 0}}
```

<br>

<ins>Usando</ins>:

```go
enderecoExemplo := endereco{"Rua dos Bobos", 0}

// colocar cada valor de uma vez
var u usuario
u.nome = "Davi"
u.idade = 21
	// {Davi 21 { 0}}

// tudo de uma vez só
usuario2 := usuario{"Davi", 21, enderecoExemplo}
	// {Davi 21 {Rua dos Bobos 0}}

// só parte dele - especifico qual parte
usuario3 := usuario{idade: 22}
	// { 22 { 0}}
```

<br>

## VII - ~~Herança~~ (só que não)

Veremos o mais perto que o go chega de herança

<ins>Declarando</ins> as structs:

```go
type pessoa struct {
	nome		string
	sobrenome	string
	idade 		uint8
	altura 		uint8
}

type estudante struct {
	pessoa
	curso		string
	faculdade	string
}
```

<br>

<ins>Usando</ins>:

```go
pessoa1 := pessoa{"Rosana", "Rezende", 34, 163}
// {Rosana Rezende 34 163}

estudante1 := estudante{pessoa1, "Engenharia", "USP"}
// {{Rosana Rezende 34 163} Engenharia USP}
// estudante1.nome => Rosana

estudante2 := estudante{{"Carlos", "Rui", 19, 180}, "Matemática", "UFES"}
```

<br>

## VIII - Ponteiros

Ponteiro é uma referência de memória
- é uma variável q salva dentro dela não um ~~valor~~, mas um <ins>endereço de memória</ins>

- sem ~~ponteiro~~

  ```go
  var variavel1 int = 10
  var variavel2 int = variavel1

	// Ambas tem o mesmo valor
  // variavel1 >> 10
	// variavel2 >> 10

	// Mas quando mudo a 2, não muda a 1 ==> era só uma cópia
  variavel2++
  // variavel1 >> 10
	// variavel2 >> 11
  ```

<br>

- com ponteiro:
	não jogamos o valor, mas o endereço de memória onde a variável tá salva

  ```go
  var variavel3 int // guarda um valor inteiro
  var ponteiro *int // guarda o endereço de memória de um valor inteiro

	// Veja q o ponteiro não sabe o valor, só a localização
  variavel3 = 100
  ponteiro = &variavel3 // pede esse & pra aceitar guardar o endereço de memória de um valor inteiro
		// variavel3 >> 100
		// ponteiro >> 0xc000016088
  ```

<br>

<ins>OBS</ins>: **Desreferenciação**: pra ver o valor na memória preciso colocar o *

```go
// *ponteiro >> 100
```

<br>

<ins>OBS</ins>: **Pq ponteiro é importante?**
- pq apesar do endereço de memória continuar sendo o mesmo, o valor do ponteiro vai mudar sempre que a variável pra qual ele aponta mude

	```go
	variavel3 = 150
		// variavel3 >> 150
		// ponteiro  >> 0xc000016088
		// *ponteiro >> 150
	```

- sempre que quero mudificar um valor sem ficar passando cópia dele posso usar ponteito

<br>

<ins>OBS2</ins>: **Valor inicial/0**: `<nil>`

```go
  var variavel3 int
  var ponteiro *int
	// variavel3 >> 0
	// ponteiro >> <nil>
```

<br>

## IX - Arrays e Slices

### Array

É uma lista de valores

<br>

<ins>Criando</ins>
- todos os ítens tem q ser do **mesmo tipo**
- tem que **especificar o tamanho**

```go
var array1 [5]string
array1[0] = "Posição 1"
// array1 >> [Posição 1    ]

array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
```

<br>

<ins>OBS</ins>: **não** consigo adicionar mais do que especifiquei no início

```go
array2[5] = "Posição 6"
// dá erro
```

<br>

<ins>OBS2</ins>: o mais perto de usar o array de forma mais fléxível é com `...`

```go
array3 := [...]int{1, 2, 3, 4, 5}
```
- fixa o tamanho do array de acordo com a quantidade de ítens que passei
- se eu passo uma posição que não existe tb dá erro
- na prática não é muito usado isso

<br>

### Slice

Estrutura baseada no array, mas **+ flexível**
- não preciso passar a quantidade
- mas todos tem que ser do mesmo tipo

Ele parece um array, mas na verdade ele tem um ponteiro que aponta pra um array

<br>

<ins>Criando</ins>

- do zero:

	```go
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	```

<br>

- de um array

  ```go
	array := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
  slice2 := array[1:3]
	// slice2 >> [Posição 2  Posição 3]
  ```

  - e se eu alterar o array? funciona **como um ponteiro**
    ```go
    array2[1] = "Posição Alterada"
    // slice2 >> [Posição Alterada  Posição 3]
    ```

<ins>Adicionando</ins>

```go
	slice = append(slice, 18)
	// slice2 >> [10 11 12 13 14 15 16 17 18]
```

<br/>

### Arrays internos

Apesar de termos visto que o slice funciona como um ponteiro de um array, tb vimos que é possível criar ele sem referenciar um array. Como isso acontece?

<br>

<ins>Criando</ins>

- podemos fazer da forma vista acima
	```go
	slice := []int{1, 2, 3}
	```

- ou usando a **função make** (ela aloca um espaço na memória pra algo... nesse caso um slice)
	- parâmetros: <ins>tipo</ins>, <ins>qtd de itens</ins>, <ins>qtd máxim (capacidade)</ins>
	- se não passar capacidade é = qtd de itens

	```go
	slice3 := make([]float32, 10, 11)
	// [0 0 0 0 0 0 0 0 0 0]
	// tamanho? len(slice3) 	>> 10
	// capacidade? cap(slice3)	>> 11
	```

<ins>Adicionando</ins>

```go
slice3 = append(slice3, 5)
// [0 0 0 0 0 0 0 0 0 0 5]
// tamanho? len(slice3) 	>> 11
// capacidade? cap(slice3)	>> 11
```

<br>

- posso ultrapassar a capacidade? **SIM**
	quando o go percebe q vai estourar, cria um novo array e dobra a capacidade dele
	```go
	slice3 = append(slice3, 6)
	// [0 0 0 0 0 0 0 0 0 0 5 6]
	// tamanho? len(slice3) 	>> 12
	// capacidade? cap(slice3)	>> 24
	```

<br>

<ins>OBS</ins>: E para **REMOVER**?

Não existe uma função específica no go... é necessário fazer na mão. #ex:

- é possível mudar o tipo do primeiro parâmetro e do retorno da função para atender o seu slice específico. Basicamente essa função vai trocar o elemento que você quer deletar com o último elemento do slice, e depois retornar um novo slice sem ele.
Porém, essa função só é recomendada se você não se importa com ordenação. 
	```go
	func remove(s []int, i int) []int {
		s[len(s)-1], s[i] = s[i], s[len(s)-1]
		return s[:len(s)-1]
	}
	```

- caso seja importante **ordenação**:
	```go
	func remove(slice []int, s int) []int {
		return append(slice[:s], slice[s+1:]...)
	}
	```

<br>

## X - Maps

Estrutura de chave valor + rígida q um struct
- chave e valor tem sempre o mesmo tipo
- a estrutura não é mutável

<br>

<ins>Criando</ins>

- map[tipo da chave]tipo do valor
<br>

  ```go
  usuario := map[string]string{
    "nome":	"Rosana",
    "sobrenome":	"Rezende",
  }
  // usuario["nome"] >> Rosana

  usuario2 := map[string]map[string]string{
    "nome": {
    	"primeiro":	"Pedro",
    	"último":	"Lucas",
    },
    "curso": {
    	"nome":		"Engenharia",
    },
  }
  // usuario2 >> map[curso:map[nome:Engenharia] nome:map[primeiro:Pedro último:Lucas]]
  ```

<br>

<ins>Deletando</ins>

```go
delete(usuario2, "nome")
// map[curso:map[nome:Engenharia]]
```

<br>

<ins>Adicionando</ins>

```go
usuario2["signo"] = map[string]string{
	"nome": "Gêmeos",
}
// map[curso:map[nome:Engenharia] signo:map[nome:Gêmeos]]
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
// Menor que 15
```

<br>

<ins>OBS</ins>: **If init**: é possível inicializar uma variável direto no if

```go
if outroNumero := -5; outroNumero > 0 {
	fmt.Println("Numero é maior que zero")
} else if numero < -10 {
	fmt.Println("Numero menor que -10")
} else {
	fmt.Println("Entre 0 e -10")
}
// Entre 0 e -10
```
- mas não consigo ~~acessar essa variável fora~~: ela está limitada ao escopo
	```go
	fmt.Println(outroNumero)
	// dá erro - diz q não tá definida
	```

<br>

## XII - Switch

Posso fazer passando no início do switch a variável a ser verificada:

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

dia := diaDaSemana(200) // Número Inválido
```

<br>

Ou passando a condição em cada caso
- é mais útil quando vc não quer avaliar a mesma variável, condições q vem de lugares diferentes

```go
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

dia2 := diaDaSemana2(1) // Domingo
```

<br>

<ins>OBS</ins>: **fallthrough**: é uma cláusula que joga o código pra dentro da <ins>próxima condição</ins>
- não é muito ~~utilizada~~

```go
func diaDaSemana3(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough 
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	...
	}
}
dia3 := diaDaSemana3(1) // Segunda-Feira
```

<br>

## XIII - Loops

O go só usa uma estrutura de repetição: **FOR**

<br>

Várias maneiras de usar:

- enquando a condição for verdadeira, faça algo (*parecido com while*)

	```go
	i := 0
	for i < 3 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	// Incrementando i
	// Incrementando i
	// Incrementando i
	// i >> 3
	```

<br>

- Inicializando a variável no escopo (*usando uma sintaxe + parecida com if init*)

	```go
	for j := 0; j < 10; j += 5 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
	// Incrementando j 0
	// Incrementando j 5
	```

<br>

<ins>Range</ins>

Em string:
- tem q transformar cada letra em string - ele acha q é o código da tabela asp

```go
for indice, letra := range "OI" {
	fmt.Println(indice, string(letra)) 
}
// 0 O
// 1 I
```

<br>

Em array ou slice:

```go
nomes := [2]string{"João", "Davi"} // podia ser slice tb
for _, nome := range nomes {
	fmt.Println(nome)
}
// João
// Davi
```

<br>

Em map:

```go

usuario := map[string]string{
	"nome":      "Leonardo",
	"sobrenome": "Silva",
}
for chave, valor := range usuario {
	fmt.Println(chave, valor)
}
// nome Leonardo
// sobrenome Silva
```

<br/>

<ins>OBS</ins>: não dá pra fazer em ~~struct~~

<br>

<ins>OBS2</ins>: infinito
- como não passou nenhuma condição é como se ela fosse sempre verdadeira `for true`

	```go
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
	```

<br>
