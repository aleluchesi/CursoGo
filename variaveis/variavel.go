package main

import (
	"fmt"
)

var (

	// Endereco deve ser Publico. Inicia com Maiusculo
	Endereco 				string 
	// Telefone deve ser publico. Inicia com Maiusculo
	Telefone		 		string
	quantidade, estoque		int   //  0
	comprou 				bool   // false
	valor 					float64  // 0.00
	palavra 				rune

)

func main() {

	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("Palavra: %v\r\n", palavra)
	fmt.Printf("valor: %s\r\n", teste)

}