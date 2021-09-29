package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	// Asignarle un valor a nuestra variable
	var horas int
	horas = 20
	fmt.Println(horas)
	fmt.Println("------------")
	// Declaración de variable corta
	minutes := 30
	fmt.Println(minutes)
	fmt.Println("------------")
	// Declaración de múltiples variables
	dato1, dato2 := "Hola",10.45
	dato3, dato4 := true, 20.1
	fmt.Println(dato1)
	fmt.Println(dato2)
	fmt.Println(dato3)
	fmt.Println(dato4)
	fmt.Println("------------")
	// Bloque de declaración de variables
	var(
		quantity 	= 25
		product 	= "Course"
		price		= 40.50
		inStock		= true
	)
	fmt.Println(quantity)
	fmt.Println(product)
	fmt.Println(price)
	fmt.Println(inStock)
	fmt.Println("------------")
	// Cuando declaramos una variable y no la utilizamos, Go no nos dejará compilar
	//y nos mostrará un error. Así, nos obliga a utilizar todas las variables que
	//estamos declarando para tener un código ordenado y limpio.

	// Constantes
	const CONSTANTE = "CONSTANTE"
	// Bloque de declaración de constantes
	const (
		PRODUCT 	= "Course"
		QUANTITY 	= 20
		PRICE		= 40.50
	)
	fmt.Println(PRODUCT)
	fmt.Println(QUANTITY)
	fmt.Println(PRICE)
	fmt.Println("------------")
	// Integers (Signed and UnSigned)
	var myVarInt int = 1
	fmt.Println(myVarInt)
	fmt.Println("------------")
	// Floats 32
	var myVarFloat32 float32 = 2
	fmt.Println(myVarFloat32)
	fmt.Println("------------")
	// Floats 64
	var myVarFloat64 float64 = 3
	fmt.Println(myVarFloat64)
	fmt.Println("------------")
	// String
	var myString string  = "Hello"
	fmt.Println(myString)
	var fname, lname string = "John", "Doe"
	fmt.Println(fname)
	fmt.Println(lname)
	fmt.Println("------------")
	// Boolean
	var myBool bool
	myBool = true
	myBool = false
	fmt.Println(myBool)
	fmt.Println("------------")
	// Byte
	var myByte byte
	fmt.Println(myByte)
	fmt.Println("------------")
	// Operadores Aritméticos
	x, y := 10, 20
	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)
	x++
	fmt.Printf("x++ = %d\n", x)
	y--
	fmt.Printf("y-- = %d\n", y)
	fmt.Println("------------")
	// Operadores de Asignación
	var r, s = 15, 25
	r = s
	fmt.Println("= ", r)
	r = 15
	r += s
	fmt.Println("+=", r)
	r = 50
	r -= s
	fmt.Println("-=", r)
	r = 2
	r *= s
	fmt.Println("*=", r)
	r = 100
	r /= s
	fmt.Println("/=", r)
	r = 40
	r %= s
	fmt.Println("%=", r)
	fmt.Println("------------")
	// Operadores de comparación
	t, u := 10, 20
	fmt.Println(t == u)
	fmt.Println(t != u)
	fmt.Println(t > u)
	fmt.Println(t >= u)
	fmt.Println(t < u)
	fmt.Println(t <= u)
	fmt.Println("------------")
	// Operadores lógicos
	var d, e, f = 10, 20, 30
	fmt.Println(d < e && d > f)
	fmt.Println(d < e || d > f)
	fmt.Println(!(d == e && d > f))
	// Operadores lógicos a nivel de bits
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0
	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Conjunción - El valor de c es %d\n", c )
	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Disyunción - El valor de c es %d\n", c )
	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Disyunción exclusiva - El valor de c es %d\n", c )
	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Corrimiento de bits a la izquierda - El valor de c es %d\n", c )
	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Corrimiento de bits a la derecha - El valor de c es %d\n", c )
	// Array Declaracion
	var array [2] string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array)
	fmt.Println("------------")
	// Slice
	var slice = []bool{true, false}
	fmt.Println(slice[0])
	var slice01 = []string{"hello","world"}
	fmt.Println(slice01)
	sliceMake := make([]int, 5) // len (a) = 5
	fmt.Println(sliceMake)
	primes := []int{2, 3, 5, 7, 11, 13}
	// Si no ponemos un valor después de los : toma hasta el fin de elementos del slice y viceversa.
	fmt.Println("------------")
	fmt.Println(primes[1:4])
	slice02 := []int{1,5}
	slice02 = append(slice02, 2, 3, 4)
	fmt.Println(slice02)
	fmt.Println("------------")
	// MAP
	var myMap01 = map[string]int{}
	myMap02 := make(map[string]string)
	fmt.Println(myMap01)
	fmt.Println(myMap02)
	fmt.Println(len(myMap01))
	var students = map[string]int{"Benjamin": 20, "Nahuel": 26}
	fmt.Println(students["Benjamin"])
	fmt.Println(students)
	students["Brenda"] = 19
	students["Marcos"] = 22
	fmt.Println(students)
	students["Benjamin"] = 22
	fmt.Println(students)
	fmt.Println(students)
	delete(students, "Benjamin")
	fmt.Println(students)
	fmt.Println("------------")
	// for
	for key, element := range students {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
	fmt.Println("------------")
	// If
	sueldo := 4500
	if sueldo > 3000 {
		fmt.Println("Esta persona debe pagar impuestos")
	}
	fmt.Println("------------")
	// If Else
	if sueldo > 3000 {
		fmt.Println("Esta persona debe pagar impuestos")
	} else {
		fmt.Println("Esta persona NO debe pagar impuestos")
	}
	fmt.Println("------------")
	// If Else If Else
	var sueldo01 float64 = 4000
	if sueldo01 <= 3000 {
		fmt.Println("Esta persona debe pagar impuestos")
	} else if sueldo01 <= 4000 {
		fmt.Printf("Debe pagar $%4.2f de su sueldo01 \n", (sueldo01/100)*10)
	} else {
		fmt.Printf("Debe pagar $%4.2f de su sueldo01 \n", (sueldo01/100)*15)
	}
	fmt.Println("------------")
	// If Corto
	if edad := 20; edad > 150 {
		fmt.Println("¿Eres inmortal?" )
	} else if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad < 18 && edad > 0 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Edad fuera del rango")
	}
	fmt.Println("------------")
	// switch
	day := 1
	switch day {
	case 0:
		fmt.Println("Lunes")
	case 1:
		fmt.Println("Martes")
	case 2:
		fmt.Println("Miércoles")
	case 3:
		fmt.Println("Jueves")
	case 4:
		fmt.Println("Viernes")
	case 5:
		fmt.Println("Sábado")
	case 6:
		fmt.Println("Domingo")
	default:
		fmt.Println("Desconocido")
	}
	fmt.Println("------------")
	// switch sin condicion
	var edad uint8 = 18
	switch {
	case edad >= 150:
		fmt.Println("¿Eres inmortal?")
	case edad >= 18:
		fmt.Println("Eres mayor de edad" )
	default:
		fmt.Println("Eres menor de edad" )
	}
	fmt.Println("------------")
	// switch con multiples casos
	today := "domingo"
	switch today {
	case "lunes", "martes", "miércoles", "jueves", "viernes":
		fmt.Printf ("%s es un día de la semana \n", today)
	default:
		fmt.Printf ("%s es un día del fin de la semana\n", today)
	}
	fmt.Println(today)
	fmt.Println("------------")
	switch day := "domigo"; day {
	case "lunes", "martes", "miércoles", "jueves", "viernes":
		fmt.Printf("%s es un día de la semana\n", day)
	default:
		fmt.Printf("%s es un día del fin de la semana \n", day)
	}
	fmt.Println(today)
	fmt.Println("------------")
	today01 := time.Now()
	var t01 int = today01.Day()
	switch t01 {
	case 5, 10, 15:
		fmt.Println("Limpia tu casa.")
	case 25, 26, 27:
		fmt.Println("Comprar comida.")
		fallthrough
	case 31:
		fmt.Println("Hoy hay fiesta.")
	default:
		fmt.Println("No hay información disponible para ese día.")
	}
	fmt.Println("------------")
	// For range
	frutas := []string{"manzana", "banana", "pera"}
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}
	fmt.Println("------------")
	// Bucle While
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
	// Romper Bucle
	fmt.Println("------------")
	sum01 := 0
	for {
		sum01++
		if sum01 >= 1000 {
			break
		}
	}
	fmt.Println(sum01)
	fmt.Println("------------")
	for i := 0; i < 10; i++{
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i, "is odd")
	}
	fmt.Println("------------")

}
