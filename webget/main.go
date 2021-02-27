package main

import (
	"io/ioutil"
	"fmt"
	"time"
	"net/http"
)

func main()  {
	
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página google", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o comando de página", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	//  chamada autenticada

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um requst", err.Error())
		return
	}
	request.SetBasicAuth("Teste", "teste")
	resposta, err = cliente.Do(request)
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
		fmt.Println("********  PASSOU AQUI *******       \r\n ")
		fmt.Println(string(corpo))
	}
	
}