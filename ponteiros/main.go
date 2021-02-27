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

	casa := new(Imovel)
	fmt.Printf("Casa é: %p -  %+v\r\n", &casa, casa)

	chacara := Imovel{10, 17, "chacara", 280000}
	fmt.Printf("Chacara é: %p -  %+v\r\n", &chacara, chacara)

	mudaimovel(&chacara)
	fmt.Printf("Chacara é: %p -  %+v\r\n", &chacara, chacara)
}

func mudaimovel(imovel *Imovel) {

	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y + 5
	
}