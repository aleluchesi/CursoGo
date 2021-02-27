package main

import (
	"fmt"
)

func main() {

	messes := 0
	situacao := true
	cidade := "Teste"

	/*  
		<   Menor
		>   Manior
		!=  Diferente
		==  Igual
		<=  Menor igual
		>=  Maior Igual
		&&  E
		||  Ou 
	*/
	if messes <= 6 {
		fmt.Println("Esse credor deve a pouco tempo!")
	}

	// Teste de Bool = True
	if situacao {
		fmt.Println("Credor está devedor!")
	}

	// Teste de Bool = False  (Colocar Exclamação "!) no incio da variavel
	if !situacao {
		fmt.Println("Credor está devedor!")
	}

	if cidade == "Teste" {
		fmt.Println("Cidade está teste")
	}

	if descricao, status := tempodevedor(messes); status {
		fmt.Println("Qual situação do cliente? ", descricao)
	}

	fmt.Println("Obrigado por nos consultar")

}

func tempodevedor(meses int) (descricao string, status bool) {

	if meses > 0 {
		status = true
		descricao = "Cliente está devendo"
		return
	}
	descricao = "Cliente não está devendo"
	return

}