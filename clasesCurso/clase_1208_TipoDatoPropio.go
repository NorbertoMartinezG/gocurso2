package clasesCurso

import (
	"fmt"
)

var a int

type dinero int

var b dinero

func TipoDeDatoEjemplo() {

	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b // no se puede asignar por a es Int y b es tipo dinero
	b = 15
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
