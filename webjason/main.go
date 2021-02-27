package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aleluchesi/curso/webjason/model"
)

func main() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um requst", err.Error())
		return
	}
	request.SetBasicAuth("Teste", "teste")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro na resposta do requst", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o comando de página", err.Error())
			return
		}
		post := model.Blogpost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o retorno JSON do servidor", err.Error())
			return
		}

		fmt.Printf("POst é: %+v\r\n", post)
	}

}
