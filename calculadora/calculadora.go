package calculadora

type Logger interface {
	Log(string) error
}

type Config interface {
	SumaEnabled(cliente string) bool
}

// Función que recibe dos enteros y retorna la suma resultante
func Sumar(num1, num2 int) int {
	return num1 + num2
}
// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Restar(num1, num2 int) int {
	return num1 - num2
}
// Función que recibe dos enteros (numerador y denominador) y retorna la división resultante
func Dividir(num, den int) int {
	return num / den
}
// Función que recibe dos enteros, un objeto del tipo logger y retorna la suma resultante
func SumarLogger(num1, num2 int, logger Logger) int {
	err := logger.Log("Ingreso a Función Sumar")
	if err != nil {
		return -99999
	}
	return num1 + num2
}

func SumarRestricted(num1, num2 int, config Config, cliente string) int {
	if !config.SumaEnabled(cliente) {
		return -99999
	}
	return num1 + num2
}