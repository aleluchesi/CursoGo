package main

import (
	"time"
	"runtime"
	"fmt"
)

func main() {

	numero := 3
	fmt.Print("O número ", numero, " se escreve assim ")
	switch numero {
	case 1:
		fmt.Printf("um.\r\n")
	case 2:
		fmt.Printf("dois.\r\n")
	case 3:
		fmt.Printf("três.\r\n")
	}


	fmt.Println("Você é da família do Unix?")
	switch runtime.GOOS {
	case "darvin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sair")
	default:
		fmt.Println("Deixa essa pergunta para lá")
	}



	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Posso dormir até mais tarde")
	default:
		fmt.Println("Vamos trabalhar!")
	}


	
	numero = 10
	fmt.Println("Quantos números cabem num digito?")
	switch  {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois digitos")
	case numero >= 100:
		fmt.Println("Número muito grande")
	}

}