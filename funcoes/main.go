package main

import (
	"cursogo/funcoes/matematica"
	"fmt"
)

func main() {

	resultado := matematica.Calcula(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da Multiplicação é: %d\r\n", resultado)

	resultado = matematica.Soma(4, 5)
	fmt.Printf("O resultado da soma: %d\r\n", resultado)

	resultado = matematica.Calcula(matematica.Divisor, 2, 2)
	fmt.Printf("O resultado da divisão é: %d\r\n", resultado)

	resultado, resto := matematica.DivisorResultado(12, 5)
	fmt.Printf("O resultado da divisão é: %d e o resto é: %d\r\n", resultado, resto)

}
