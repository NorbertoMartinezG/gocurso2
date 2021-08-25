package clasesCurso

import "fmt"

func Clase1203() {
	// ...interface{} ... cantidad variable de datos
	fmt.Println("clase ", 1203, " hola")

	// Operadores
	/*
			Operador de declaración corta
		●	Terminología
			○	Palabras claves
				■	Son palabras reservadas para uso interno de Go
					●	Algunas veces se les llama  “palabras reservadas”
					●	No puedes usar una palabra clave para algo distinto a lo que está destinada.
		○	Operador
			■	En “2 + 4” el “+” es el OPERADOR
			■	Un operador es un caracter que representa una acción, como por ejemplo “+” es un OPERADOR aritmético que representa adición.

		○	operando
			■	En “2 + 2” los “2” son OPERANDOS
				○	Declaración
					■	En programación una declaración es el elemento más pequeño de un programa que expresa una acción que va a ser llevada a cabo. Es una instrucción que la cual le da el comando a la computadora para ejecutar una acción específica. Un programa es formado por una secuencia de una o más declaraciones.
				○	Expresión
					■	En programación una expresión es una combinación de uno o más valores explícitos, constantes, variables, operadores y funciones que el lenguaje de programación interpreta y computa para producir otro valor. For ejemplo, 2+3 es una expresión el cual evalúa a 5.


	*/
	x := 42
	y := 2 + 4 // y = declaracion, 2 + 4 = expresion 2 = operando
	fmt.Println(x, y)

}
