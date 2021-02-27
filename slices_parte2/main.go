package main

import (
	"fmt"
)

func main()  {
	
	cidade := make([]string, 5)
	cidade[0] = "New York"
	cidade[1] = "São Paulo"
	cidade[2] = "Madeira"
	cidade[3] = "Tokio"
	cidade[4] = "Cingapura"

	fmt.Println(cidade, len(cidade), cap(cidade))	

	for indice, cidade := range cidade {
		fmt.Printf("Cidade [%d]  =  %s \n\r", indice, cidade)
	}

	// Primeiro item começa com indice 0
	// Segundo item começa com indice 1

	captaisaisa := cidade[3:5]
	fmt.Println(captaisaisa, len(captaisaisa), cap(captaisaisa))

	// Mostra do primeeiro até o desejado. No caso 2
	temp1 := cidade[:2]
	fmt.Println(temp1, len(temp1), cap(temp1))

	// Mostra do desejado até o ultimo. No caso os 2 últimos.
	temp2 := cidade[len(cidade)-2:]
	fmt.Println(temp2, len(temp2), cap(temp2))

	// remover um item do slice. Criar um slice temporário sem o item que deseja remover.
	// cria um slice ate o item desejado excluir.
	// fazer um Append dos item após o item deseja excluir 
	// copia para o slice inical

	indiceRetirar := 2

	temp := cidade[:indiceRetirar]
	temp = append(temp, cidade[indiceRetirar+1:]...)

	copy(cidade, temp)

	fmt.Println(cidade, len(cidade), cap(cidade))	

	for indice, cidade := range cidade {
		fmt.Printf("Cidade [%d]  =  %s \n\r", indice, cidade)
	}

}