package matematica

func Calcula(funcao func(int, int) int, numero1, numero2 int) int {

	return funcao(numero1, numero2)

}

func Multiplicador(x, y int) int {
	
		return x * y
	
	}

func Divisor(x, y int) (resultado int) {

	resultado = x / y
	return

}

func DivisorResultado(x int, y int) (resultado int , resto int) {

	resultado = x / y
	resto = x % y
	return

}