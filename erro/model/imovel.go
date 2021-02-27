package model

import "errors"

type Imovel struct {

	X	  int 		`json:"Cordenada_X"`
	Y	  int 		`json:"Cordenda_Y"`
	Nome  string 	`json:"nome"`
	valor int
}

var ErrValorInvalido = errors.New("Valor não válido!")
var ErrValorMaior = errors.New("O valor informado é muito alto!")

// Atribui valor ao imóvel
func (i *Imovel) SetValor (valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorInvalido
		return
	} else if valor > 10000000{
		err = ErrValorMaior
		return
	}
	i.valor = valor
	return
}

// Carrega valor do imóvel 
func (i *Imovel) GetValor() int {
	return i.valor
}