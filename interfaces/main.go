package main

import "github.com/aleluchesi/cursodego/interfaces/model"
import "fmt"

func main()  {
	
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	QueroAcordarCacarejo(jojo)
	QueroAcordarQuack(jojo)

}

func QueroAcordarCacarejo(g model.Galinha){
	fmt.Println(g.Cacareja())
}

func QueroAcordarQuack(p model.Pato) {
	fmt.Println(p.Quack())
}