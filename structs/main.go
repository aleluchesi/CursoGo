package main

import (
	"fmt"
)

type Imovel struct{
	X 		int
	Y 		int
	Nome 	string
	valor 	int
}

func main() {

	casa := Imovel{}  //Inicialização vazia
	fmt.Printf("Casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 50, "meu apto", 760000}
	fmt.Printf("Apto é: %+v\r\n", apartamento)

	chacara := Imovel {
		Y: 70,
		Nome: "Chacara",
		valor: 850000,
		X: 10,
	}
	fmt.Printf("Chacara é: %+v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	casa.valor = 50000
	casa.X = 18
	casa.Y = 20
	fmt.Printf("Casa é: %+v\r\n", casa)


}