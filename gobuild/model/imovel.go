package model

type Imovel struct {

	X	  int 		`json:"Cordenada_X"`
	Y	  int 		`json:"Cordenda_Y"`
	Nome  string 	`json:"nome"`
	valor int
}


// Atribui valor ao imóvel
func (i *Imovel) SetValor (valor int) {
	i.valor = valor
}

// Carrega valor do imóvel 
func (i *Imovel) GetValor() int {
	return i.valor
}