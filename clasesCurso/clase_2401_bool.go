package clasesCurso

import "fmt"

func Clase2401() {
	/*
		Fundamentos de Programación
			Tipo Booleano
			1.	Un booleano es un TIPO binario el cual puede tener dos posibles valores “verdadero” y “falso”
			2.	Cuando usas un operador de comparación de igualdad, este es una expresión el cual evalúa a un valor booleano
				a.	==
				b.	<=
				c.	>=
				d.	!=
				e.	<
				f.	>
	*/

	var x bool
	fmt.Println(x)
	x = true
	fmt.Println(x)

	x1 := 7
	y1 := 15

	fmt.Println("x1 es igual a y1")
	fmt.Println(x1 == y1)
	fmt.Println("x1 es mayor o igual a y1")
	fmt.Println(x1 >= y1)
	fmt.Println("x1 es menor o igual a y1")
	fmt.Println(x1 <= y1)
	fmt.Println("x1 es diferente a y1")
	fmt.Println(x1 != y1)

}
