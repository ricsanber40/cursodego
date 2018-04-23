package matematica

//Calculo é uma funçao
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador mutiplica
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor divide
func Divisor(x int, y int) (resultado int) {
	resultado = x / y
	return
}

//DivisorComResto retorna 2 valores
func DivisorComResto(x int, y int )(resultado int, resto int){
	resultado = x / y
	resto = x % y
	return
}