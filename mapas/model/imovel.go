package model

type Imovel struct {

	X	  int 		
	Y	  int 		
	Nome  string 	
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