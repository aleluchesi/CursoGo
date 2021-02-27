package model

type Galinha interface {
	Cacareja() string
}

type Pato interface {
	Quack() string 
}

type Ave struct {
	Nome string
}

func (a Ave) Cacareja() string {
	return "Cacareja....."
}

func (a Ave) Quack() string {
	return "Quack Quack..."
}