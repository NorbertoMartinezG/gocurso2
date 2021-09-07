package clasesCurso

import "fmt"

func Clase1305() {
	//Ejercicios - Ninja Nivel 1
	fmt.Println("Ejercicio 4 - Ninja Nivel 1")
	/*
		Ejercicio práctico #4
			●	Información - documentación interesante y nueva terminología “tipo subyacente, raíz o tipo implícito”
			○	https://golang.org/ref/spec#Types
			Para este ejercicio
			1.	Crea tu propio tipo. Haz que el tipo subyacente, raíz o implícito sea un int.
			2.	Crea una VARIABLE de tu nuevo TIPO con el IDENTIFICADOR “x” usando la palabra clave “VAR”
			3.	En func main
				a.	Imprime el valor de la variable “x”
				b.	Imprime el tipo de la variable “x”
				c.	Asigna 42 a la VARIABLE “x” usando el OPERADOR “=”
				d.	Imprime el valor de la variable “x”
	*/
	type numero int // tipo subyacente INT

	var x numero // variable tipo numero

	fmt.Println(x)
	fmt.Printf("El tipo de x es: %T\n", x) //%T mapea el tipo de dato, muestra el tipo de dato x (clasesCurso.numero)

	x = 42
	fmt.Println(x)
	fmt.Println("")

	/*
		Ejercicio práctico #5
			A partir del código del ejercicio anterior
			1.	A nivel de paquete usando la palabra clave “var”, crear una VARIABLE con el IDENTIFICADOR “y”.
				La variable debería ser del mismo TIPO SUBYACENTE que tu TIPO “x” creado anteriormente
			a.	ejemplo:
				type hotdog int
				var x hotdog
				var y int

			2.	en func main
				a.	Esto lo debería tener listo
					i.	Imprimir el valor de la variable “x”
					ii.	Imprimir el tipo de la variable “x”
					iii.	Asigna un VALOR a la VARIABLE “x” usando el OPERADOR “=”
					iv.	Imprime el valor de la variable “x”
				b.	Ahora haces esto
					i.	Ahora usa CONVERSIÓN para convertir el TIPO del VALOR almacenado en “x” al TIPO IMPLÍCITO
						1.	Usa el operador “=” para ASIGNAR ese valor a “y”
						2.	Imprime el valor almacenado en “y”
						3.	Imprime el tipo de “y”
	*/
	fmt.Println("Ejercicio 5 - Ninja Nivel 1")
	var x1 numero
	var y int

	fmt.Println(x1)
	fmt.Printf("El tipo de x es: %T\n", x1) // imprimir tipo

	x1 = 42
	fmt.Println(x1)

	y = int(x1)         // conversion tipo "numero" a int
	fmt.Println(y)      // valor de y
	fmt.Printf("%T", y) // tipo de y

}
