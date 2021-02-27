package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"time"
	"net/http"
	"github.com/aleluchesi/cursodego/webpost/model"
)

func main()  {
	
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Alessandro Luchesi"

	conteudoenviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao gerar JSON", err.Error())
		return
	}	

	request, err := http.NewRequest("POST", "https://requestb.in/1jd32x91", bytes.NewBuffer(conteudoenviar))
	if err != nil {
		fmt.Println("[main] Erro ao post um requst", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao chamar o serviço", err.Error())
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
	
}