package main

import (
	"sync"
	"encoding/json"
	"strings"
	"encoding/csv"
	"bufio"
	"fmt"
	"os"
	"time"
	"github.com/aleluchesi/cursodego/goroutines/model"
)

var farol sync.WaitGroup

func main()  {

	farol.Add(2)
		go traduzarquivoJSON("saopaulo")
		go traduzarquivoJSON("riodejaneiro")
	farol.Wait()

}

func traduzarquivoJSON(nomearquivo string) {

	fmt.Println(time.Now(), "  *** Começando a traduzir ***")
	arquivo, err := os.Open(nomearquivo + ".csv")
	if err != nil {
		fmt.Println("[Main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}
	defer arquivo.Close()

	leitorCSV := csv.NewReader(arquivo)
	conteudo, err := leitorCSV.ReadAll()
	if err != nil {
		fmt.Println("[Main] Houve um erro ao ler o arquivo. Erro: ", err.Error())
		return	
	}

	arquvioJSON, err := os.Create(nomearquivo + ".json")
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

	fmt.Println(time.Now(), "  *** Finalizando a tradução ***")
	farol.Done()

}