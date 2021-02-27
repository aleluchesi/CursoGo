package main

import (
	"fmt"
	"cursodego/variaveis/pacotes/prefixo"
	"cursodego/variaveis/pacotes/operadora"
)

var NomeUsuario = "Alessandro"

func main() {

	fmt.Printf("Nome do Usuário: %s\r\n", NomeUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome Operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("NOme completo: %s\r\n", operadora.NomeCapital)

}