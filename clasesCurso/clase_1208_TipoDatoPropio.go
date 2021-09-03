package clasesCurso

import "fmt"

var a int

//type dinero int

var b int

func TipoDatoPropio() {

	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a) // %T muestra el tipo de dato

	b = b + 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = b
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
