package clasesCurso

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func Clase4101() {

	multi_line := `Struct
	Un Struct es un tipo de dato compuesto. (tipos de datos compuestos, también, tipos de datos 
	agregados o tipos de datos complejos). Los structs nos permiten componer valores con otros 
	valores de diferentes tipos.
	
		`
	fmt.Printf("%s", multi_line)

	p1 := persona{
		nombre:   "Eduar",
		apellido: "Tua",
		edad:     25,
	}

	p2 := persona{
		nombre:   "Cóndor",
		apellido: "Pérez",
		edad:     50,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.nombre, p1.apellido, p1.edad)
	fmt.Println(p2.nombre, p2.apellido, p2.edad)

}

type agenteSecreto struct {
	persona
	lpm bool
}

func Clase4102() {

	multi_line := `Structs embebidos
	Podemos tomar un struct y embeberlo en otro struct. 
	Cuando haces esto el tipo interno es promovido al tipo externo.	
		`
	fmt.Printf("%s", multi_line)

	as1 := agenteSecreto{persona: persona{

		nombre:   "James",
		apellido: "Bond",
		edad:     31,
	}, lpm: true,
	}

	p2 := persona{
		nombre:   "Condor",
		apellido: "Pérez",
		edad:     55,
	}

	fmt.Println(as1)
	fmt.Println(p2)
	fmt.Println(as1.nombre, as1.apellido, as1.edad)
	fmt.Println(p2.nombre, p2.apellido, p2.edad)

}

var x, y int

func Clase4103() {

	multi_line := `Es bueno familiarizarse con el lenguaje usado para hablar del lenguaje Go. 
	Las “especificaciones del lenguaje” pueden resultar difíciles de leer. 
		
		`
	fmt.Printf("%s", multi_line)

	x = 42
	y = 43
	fmt.Println(x, y)

	x, y = 44, 45
	fmt.Println(x, y)

}

func Clase4104() {

	multi_line := `Un struct anónimo es un struct el cual no es asociado con un 
	identificador en específico. 
		
		`
	fmt.Printf("%s", multi_line)

	p := struct {
		nombre, apellido string
		edad             int
	}{
		nombre:   "Eduar",
		apellido: "Tua",
		edad:     25,
	}

	fmt.Println(p)

}

//Creamos VALORES de ciertos TIPOS que se almacenan en VARIABLES
//las VARIABLES tienen IDENTIFICADORES

type persona2 struct {
	nombre   string
	apellido string
}

type fooo int

var y22 fooo

func Clase4105() {

	multi_line := `Es bueno familiarizarse con el lenguaje usado para hablar del lenguaje Go. 
	Las “especificaciones del lenguaje” pueden resultar difíciles de leer. 
		
		`

	fmt.Printf("%s", multi_line)

	//var x int

	y22 = 42
	p1 := persona2{
		nombre:   "Eduar",
		apellido: "Tua",
	}
	fmt.Println(p1)
	fmt.Printf("%T", y22)

}
