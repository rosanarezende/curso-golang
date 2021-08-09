# 01 - FUNÇÕES AVANÇADAS

<br/>

## 1. Retorno Nomeado

```go
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
soma, subtracao := calculosMatematicos(10, 20)
// 30 -10
```

<br>

## 2. Funções Variáticas

Pode receber n parâmetros.
- os parâmetros formam um slice e podemos iterar nele


```go
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

totalDaSoma := soma(1, 2, 3, 0) // 6
```

<br>

```go
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

escrever("Olá Mundo", 10, 2, 3)
// Olá Mundo 10
// Olá Mundo 2
// Olá Mundo 3
```

<br>

<ins>OBS</ins>: só dá pra ter 1 parâmetro variático por função... e tem que ser o último.

```go
func escrever2(texto ...string, numeros ...int) {
	...
}
// dá erro
```

<br>

## 3. Funções Anônimas

Pode ser com ou sem parâmetro

E posso ou não armazenar seu retorno dentro de uma variável

<br>

```go
func() {
	fmt.Println("Olá Mundo")
}()


retorno := func(texto string) string {
	return fmt.Sprintf("Recebido -> %s", texto)
}("Passando Parâmetro")

fmt.Println(retorno) // Recebido -> Passando Parâmetro
```

<br>

## 4. Funções Recursivas

São funções que chamam elas mesmas

- #ex: uma função q retorna num número de determinada posição na sequencia de fibonacci (*1 1 2 3 5 8 13 21 34 ...*)

	```go
	func fibonacci(posicao uint) uint {
		if posicao <= 1 {
			return posicao
		}
		return fibonacci(posicao-2) + fibonacci(posicao-1)
	}

	fmt.Println(fibonacci(4)) // 3

	posicao := uint(3)
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	// 1
	// 1
	// 3
	```

<br>

<ins>OBS</ins>: Não deve ser usada quando temos ~~números muito grandes~~, pq ela vai ficar parada fazendo as operações e pode gerar, além da demora, um "estouro de pilha" (stack overflow)... q é um erro por empilhar muitas execuções e ficar muito pesado pro processamento.

<br>

## 5. Defer

Uma cláusula que adia a execução de determinado pedaço de código até o último momento possível
- se main, até o final da função
- se forem funções com retorno, imediatamente antes do retorno

<br>

- ex1: main

	```go
	func funcao1() {
		fmt.Println("Executando a função 1")
	}
	func funcao2() {
		fmt.Println("Executando a função 2")
	}

	func main() {
		defer funcao1()
		funcao2()
		// Executando a função 2
		// Executando a função 1
	}
	```
	- adiou até main acabar

<br>

- #ex2: função com retorno

	```go
	func alunoEstaAprovado(n1, n2 float32) bool {
		defer fmt.Println("Média calculada. Resultado será retornado!")
		fmt.Println("Entrando na função para verificar se o aluno está aprovado")

		media := (n1 + n2) / 2
		if media >= 6 {
			return true
		}
		return false
	}

	fmt.Println(alunoEstaAprovado(7, 8))
	// Entrando na função para verificar se o aluno está aprovado
	// Média calculada. Resultado será retornado!
	// true
	```
	- qual a utilidade do defer nesse caso?
	se não teríamos que colocar a msg antes do return true e antes do return false, gerando repetição de código

<br>

O defer é muito útil quando estamos lidando com BD. 
- #ex: 
	- se der tudo certo, quero retornar uma mensagem de ok e fechar a conexão com o BD
	- se no meio da operação ocorre um erro e precisamos dar um return antes do esperado... independente disso queremos fechar a conexão com o BD
	
	Nesse caso é melhor abrir a conexão e em seguida usar o defer pra fechar, independente do que aconteça

<br>

## 6. Panic e Recover

Muito utilizadas com o defer


- ex: tenho uma função que calcula média
	- **PANIC**: por algum motivo se a média for 6 o programa entra em pânico e não sabe o que fazer
		- <ins> a função panic interrompe o fluxo normal do programa e para tudo </ins>
		- não é um ~~erro~~
	- **RECOVER**: 
		- <ins>é uma maneira de recuperar um programa que está entrando em pânico</ins>
		- se não houver pânico o programa ignora a função (pq é nil)

```go
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(8, 6)) 
	fmt.Println("Pós execução!") 
	// true
	// Pós execução!


	fmt.Println(alunoEstaAprovado(8, 6)) 
	fmt.Println("Pós execução!")

	// ==== se só tivesse o panic
		// panic: A MÉDIA É EXATAMENTE 6!
		// ... (msg de erro)


	// ==== panic + recover
		// Execução recuperada com sucesso!
		// false
		// Pós execução!
			// nesse caso deu false pq ele retornou o valor 0 do booleano
}
```

<br>

## 7. Closure

Funções que referenciam variáveis q estão fora do seu corpo

```go
func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto) // Dentro da função main

	funcaoNova := closure()
	funcaoNova() // Dentro da função closure
}
```
- veja que a funcaoNova não perdeu a referência do texto

<br>

## 8. Ponteiros

Vamos comparar:

- função comum (sem ~~ponteiro~~)

	```go
	func inverterSinal(numero int) int {
		return numero * -1
	}

	numero := 20
	numeroInvertido := inverterSinal(numero)
	// numeroInvertido >>	-20
	// numero >> 		20
	```
	- como o que passei é uma cópia do número, ele não muda
	- <ins>uso</ins>: quando não quiser ~~alterar o valor original~~

	<br>

- **com ponteiro**: 
	```go
	func inverterSinalComPonteiro(numero *int) {
		*numero = *numero * -1
	}

	numero := 20
	inverterSinalComPonteiro(&numero)
	// numero >> -20
	```
	- afeta o número (não é ~~cópia~~, mas um endereço de memória)
	- <ins>uso</ins>: quando quiser **alterar o valor original**

<br>

## 9. Init

É uma função que será **executada antes da main**.

```go
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)

	// Executando a função init
	// Função main sendo executada
}
```
- a ordem não influencia

<br>

<ins>OBS</ins>: podemos ter **uma por arquivo**, não ~~uma por pacote~~.

<br>


# 02 - Métodos

São parecidos com funções, mas está obrigaoriamente **associado a alguma coisa** (struct, interface)

<br>

- <ins>ex</ins>: vamos criar um usuário (struct) e um método que toda vez que invocado salve ele no BD

	```go
	type usuario struct {
		nome  string
		idade uint8
	}

	func (u usuario) salvar() {
		fmt.Printf("Salvando os dados de %s no BD.\n", u.nome)
	}
	usuario2 := usuario{"Usuario 2", 30}
	usuario2.salvar()
	// Salvando os dados de Usuario 2 no BD
	```
	- relembrando:
		- `%s` - string
		- `%d` - números
		- `%f` - número real

<br>

- agora criaremos mais métodos pra esse usuários... saber se ele é maior e aumentar sua idade (alterar um parêmetro dele)

	```go
	func (u usuario) maiorDeIdade() bool {
		return u.idade >= 18
	}
	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println("É maior de idade? ", maiorDeIdade)
	// É maior de idade?  true
	```

<br>

- tb é possível **alterar um atributo** desse usuário - ex: aumentar sua idade
	- nesse caso usamos <ins>ponteiro</ins>

	```go
	func (u *usuario) fazerAniversario() {
		u.idade++
	}
	usuario2.fazerAniversario()
	fmt.Println("Nova idade: ", usuario2.idade)
	// Nova idade:  31
	```

<br>

# 03 - Interfaces

Muito usadas quando precisamos ter flexibilidade com tipos.

A interface só tem assinatura de métodos

E sua implementação é implícita... o go entende

<br>

Vejamos um exemplo com **FORMAS**:

- primeiro crio uma interface forma, que tem q ter um método área retornando um float64

	```go
	type forma interface {
		area() float64
	}
	```

- depois quero criar a forma RETÂNGULO

	```go
	// faço a struct
	type retangulo struct {
		altura  float64
		largura float64
	}
	// implemento o método área -- ele passa a ser forma
	func (r retangulo) area() float64 {
		return r.altura * r.largura
	}
	```

- e tb quero criar a forma CÍRCULO

	```go
	type circulo struct {
		raio float64
	}
	// tb implemento o método área - agora é forma
	func (c circulo) area() float64 {
		return math.Pi * math.Pow(c.raio, 2)
	}
	```

- agora podemos escrever uma função, q ao invés de receber uma struct recebe uma interface, bem mais flexível
	```go
	func escreverArea(f forma) {
		fmt.Printf("A área da forma é %0.2f\n", f.area())
	}

	r := retangulo{10, 15}	
	escreverArea(r) // A área da forma é 150.00

	c := circulo{10}
	escreverArea(c) // A área da forma é 314.16
	```

<br>

### TIPO GENÉRICO

É uma variação da interface que podemos usar, mas com cuidado pra não virar uma gambiarra.
- podemos usar uma interface como sendo um tipo genérico pra função aceitar qq coisa

<br>

- #ex: uma função que imprime qq tipo de dado q passemos pra ela

	```go
	func imprimeDeTudo(interf interface{}) {
		fmt.Println(interf)
	}

	imprimeDeTudo("String") // String
	imprimeDeTudo(1) // 1
	imprimeDeTudo(true) // true
	```
	- um exemplo disso é o Println, que recebe uma interface de qq tipo

<br>

- ex: um map que aceita qq dado como chave e valor

	```go
	mapeiaDeTudo := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	// mapeiaDeTudo >> map[true:12 100:true 1:String String:String]
	```
	- isso pode nos gerar muita bagunça

<br>
