# 01 - FUNÇÕES AVANÇADAS

## 1. Retorno Nomeado

```go
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
soma, subtracao := calculosMatematicos(10, 20)
```

<br>

## 2. Funções Variáticas

```go
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

totalDaSoma := soma(1, 2, 3, 0)
//
```

<br>

```go
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7)
//
```

<br>

## 3. Funções Anônimas

```go
  retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
  //
```

<br>

## 4. Funções Recursivas

```go
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

posicao := uint(12)
for i := uint(1); i <= posicao; i++ {
	fmt.Println(fibonacci(i))
}
```

<br>

## 5. Defer

```go
func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

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
//
```

<br>

## 6. Panic e Recover

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
}
```

<br>

## 7. Closure

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
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
```

<br>

## 8. Ponteiros

```go
func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

}
```

<br>

## 9. Init

```go
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
```

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

## FORMAS

```go

```

<br>

## TIPO GENÉRICO

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
