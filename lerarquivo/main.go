package main

import (
	"encoding/csv"
	//"bufio"
	"fmt"
	"os"
)

func main()  {
	
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[Main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	/*  arquivo TXT
		scanner :=  bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linha := scanner.Text()
			println("O conteúdo da linha é: ", linha)
		}
	*/

	// Trabalhar com arquivo CSV

	leitorCSV := csv.NewReader(arquivo)
	conteudo, err := leitorCSV.ReadAll()
	if err != nil {
		fmt.Println("[Main] Houve um erro ao ler o arquivo. Erro: ", err.Error())
		return	
	}
	for indiceLinha, linha := range conteudo {
		fmt.Printf("Linha [%d] é %s\r\n", indiceLinha, linha)
		for indiceitem, item := range linha {
			fmt.Printf("Item [%d] é %s\r\n", indiceitem, item)
		}
	}

	arquivo.Close()

}