# RESUMO DAS AULAS

### I. Vantagens de prender Go (Golang)

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

### II - Anotações



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
