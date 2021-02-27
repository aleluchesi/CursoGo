package main

import "github.com/aleluchesi/cursodego/structs_avancado/model"
import "fmt"
import "encoding/json"

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 10
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("Casa é %+v\r\n", casa)
	fmt.Printf("Valoe é: %d\r\n", casa.GetValor())

	objJASON, _ := json.Marshal(casa)

	fmt.Println("Casa em Json: ", string(objJASON))


}