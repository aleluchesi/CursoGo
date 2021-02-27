package main

import "github.com/aleluchesi/cursodego/erro/model"
import "fmt"
import "encoding/json"

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 10
	casa.Y = 25
	
	if err := casa.SetValor(100000000); err != nil {
		fmt.Println("[main] - ", err.Error())
		if err == model.ErrValorMaior {
			fmt.Println("[Main] Verificar com corretor")
		}
		return
	}

	fmt.Printf("Casa é %+v\r\n", casa)
	fmt.Printf("Valoe é: %d\r\n", casa.GetValor())

	objJASON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("[main] houve erro na geração do objeto Json", err.Error())
		return
	}
	fmt.Println("Casa em Json: ", string(objJASON))


}