package clasesCurso

import "fmt"

var a1 int // NO SE PERMITE VOLVER A DEFINIR UNA VARIABLE CON EL MISMO NOMBRE QUE YA SE DEFINIO EN OTRO ARCHIVO DEL PROYECTO

//type dinero int NO ES NECESARIO VOLVER A DECLARARLO CUANDO YA SE DECLARO EN OTRO ARCHIVO

var b1 dinero

func ConversionDatos() {
	/*
		Conversión, no casting

			Go tiene un lenguaje para hablar de su lenguaje. Algunos términos han sido descartados porque traen asociados conceptos diferentes.
			Una nueva mirada al futuro de la programación. Nuevas palabras para hablar de algunos conceptos. No hablamos de objetos, en Go hablamos
			de crear TIPOS y VALORES de ciertos TIPOS. No hablamos de casting, hablamos de CONVERSIÓN y ASERCIÓN.
	*/

	a1 = 42
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	b1 = 1000
	fmt.Println(b1)
	fmt.Printf("%T\n", b1)

	a1 = int(b1) // conversion de dato dinero a int.
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

}
