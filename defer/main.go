package main

import (
	"encoding/json"
	"strings"
	"encoding/csv"
	"bufio"
	"fmt"
	"os"
	"github.com/aleluchesi/cursodego/escreverarquivos/model"
)

func main()  {
	
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[Main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}
	defer arquivo.Close()

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

	arquvioJSON, err := os.Create("cidade.json")
	if err != nil {
		fmt.Println("[Main] Houve um erro ao criar o arquivo. Erro: ", err.Error())
		return	
	}
	defer arquvioJSON.Close()

	escritor := bufio.NewWriter(arquvioJSON)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for indiceitem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)
			if err != nil {
				fmt.Println("[Main] Houve um erro no item ", item , " Erro: ", err.Error())
				return	
			}	
			escritor.WriteString("  " + string(cidadeJSON))
			if indiceitem + 1 < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	escritor.WriteString("\r\n]")
	
	//  Arquivo bufio trabalhado em memória. Precisa dar o Flush para tirar da memória e carrgar no arquivo físico 
	escritor.Flush()


}