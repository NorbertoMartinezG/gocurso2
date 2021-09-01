package clasesCurso

import "fmt"

// Declarando la variable con el identificador z de tipo int
var z = 41
var variable int

func Clase1204() {

	/*
			La palabra clave var
		●	Parentesis
			( )
		●	Llaves
			{ }
		●	Dónde puede ser usada var
			○	Cualquier lugar dentro del paquete
		●	scope
			○	Dónde una variable existe y es accesible
			○	Buena práctica: mantén el scope lo más “reducido” posible

			Explorando tipos de datos
			●	DECLARAR una VARIABLE de un cierto TIPO, sólo puede contener VALORES de ese TIPO
			●	“Go suffers no fools.”
				○	like “dead men tell no tales”
			●	var z int = 21
				○	package scope
			●	Tipos de datos primitivos
				○	En ciencias de computación, un tipo de datos primitivo es uno de los siguientes:
					■	Un tipo de datos básico es un tipo de datos que proporcionado por un lenguaje de programación como un bloque básico de construcción de código. La mayoría de los lenguajes permiten recursivamente construir tipos de datos compuestos más complejos, utilizando como inicio tipos de datos básicos.
					■	Un tipo de datos interno es un tipo de datos al cual el lenguaje proporciona soporte interno.
				○	En la mayoría de los lenguajes de programación, todos los tipos de datos básicos son de incorporación interna del lenguaje. Además, muchos lenguajes también proporcionan un conjunto de tipos de datos compuestos. Hay varias opiniones de si un tipo de dato incorporado internamente en el lenguaje que no es básico debería ser considerado “primitivo”.
			●	Tipos de datos compuestos
				○	En ciencias de computación,un tipo de datos compuesto es cualquier tipo de dato el cual puede ser construido en un programa usando los tipos de datos internos del lenguaje de programación u otros tipos de datos compuestos. Algunas veces es llamado structure o aggregate data type, aunque el último término puede referirse a arreglos, listas, entre otros. El acto de construir un tipo de datos compuesto es conocido como  composición

	*/
	// variables
	var w int
	x := 42 + 7
	y := "James bond"
	fmt.Println(x)
	fmt.Println(y)
	x = 50
	fmt.Println(x)
	fmt.Println(z)
	fmt.Println(w)

	// tipo de datos
	variable = 55
	fmt.Println(variable)

}
