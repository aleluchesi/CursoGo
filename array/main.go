package main

import (
	"fmt"
)

func main()  {
	
	var teste[3] int

	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Println("Qual a capacidade do array? ", len(teste))

	cantores := [2]string{"Michael Jackson", "John Lenon"}

	fmt.Printf("O que há no Array? \r\r%v\r\n", cantores)

	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "New Yourk"}
	for indice, cidade := range capitais {
		fmt.Printf("A capital [%d] é %s\n\r", indice, cidade )
	}

}