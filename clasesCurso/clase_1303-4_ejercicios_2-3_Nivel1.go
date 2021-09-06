package clasesCurso

import "fmt"

var x3 int
var y3 string
var z3 bool

func Clase1303() {
	//Ejercicios - Ninja Nivel 1
	fmt.Println("Ejercicios - Ninja Nivel 1")
	/*	--caracteres simbolos de escape
		\a   U+0007 alert or bell
		\b   U+0008 backspace
		\f   U+000C form feed
		\n   U+000A line feed or newline
		\r   U+000D carriage return
		\t   U+0009 horizontal tab
		\v   U+000B vertical tab
		\\   U+005C backslash
		\'   U+0027 single quote  (valid escape only within rune literals)
		\"   U+0022 double quote  (valid escape only within string literals)

		Ejercicio práctico #2
		1.	Usa var para DECLARAR tres VARIABLES. Las variables deben tener scope a nivel de paquete.
			No asignar VALORES a las variables. Usa los siguientes IDENTIFICADORES para las variables y
			asegúrate de que las variables son de los siguientes TIPOS (lo quiere decir que pueden
			almacenar VALORES de ese TIPO).
				a.	identificador “x” tipo int
				b.	identificador “y” tipo string
				c.	identificador “z” tipo bool
		2.	En main
				a.	Imprime los valores de cada identificador
				b.	El compilador asigna valores a las variables. ¿Cómo son llamados esos valores?
	*/
	fmt.Println("ejercicios 1303")
	fmt.Println("el compilador asigna valor 0 a variables no inicializadas")
	fmt.Println(x3)
	fmt.Println(y3)
	fmt.Println(z3)

	/*
		Ejercicio práctico #3
		Usando el código del ejercicio anterior,
		1.	En scope a nivel de paquete, asigna los siguientes valores a las tres variables
			a.	a x asignale 42
			b.	a y asignale “James Bond”
			c.	a z asignale true
		2.	en main
			a.	Usa fmt.Sprintf para imprimir todos los VALORES en un solo string. ASIGNA el valor retornado de TIPO string usando el operador de declaración corta a  la VARIABLE con el IDENTIFICADOR “s”
			b.	Imprime el valor almacenado por la variable “s”
	*/
	x3 = 42
	y3 = "paquito"
	z3 = true
	fmt.Println("ejercicios 1304")
	s := fmt.Sprintf("%v\t %v\t %v\t", x3, y3, z3)
	fmt.Println(s)

}
