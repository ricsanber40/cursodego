package main

import (
	"fmt"
	"projetos/cursodego/pacotes/operadora"
	"projetos/cursodego/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do Usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("TesteComPrefixo da Capital: %s\r\n", prefixo.TesteComPrefixo)
}
