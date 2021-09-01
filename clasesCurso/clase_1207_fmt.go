package clasesCurso

import (
	"fmt"
	"io"
	"os"
)

func FmtFuncion() {
	fmt.Println("usos de paquete fmt")

	// El paquete fmt
	// ●	Configuración básica del código
	// ○	usando var
	// ■	using zero value
	// ○	Usando el operador de declaración corta
	var a int
	var b string = "Programa"
	var c bool
	var d bool = true

	e := 42
	f := "dice hola mundo"
	g := `El doctor dice que comer vegetales es saludable`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("Mi", b, f) // cantidad variable de parametros Println respeta espacios entre parametros al imprimir
	fmt.Println(e)
	fmt.Println(g)

	//----------------------------
	var a1 int
	var b1 string = " Programa "

	fmt.Println(a1)
	fmt.Println(b1)
	/*
		The default format for %v is:

		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	*/
	// func Printf
	// func Printf(format string, a ...interface{}) (n int, err error)

	fmt.Printf("El valor de la variable a1 es: %d\n", a1)
	fmt.Printf("El valor de la variable b1 es: %s\n", b1)
	fmt.Printf("El tipo de a1 es: %T\n", a1)
	fmt.Printf("El tipo de b1 es: %T\n", b1)

	//-----------------------------

	var a3 int
	var b3 string = "James Bond"
	//var c bool
	//var d bool = true

	//e := 42
	//f := "Shaken not stirred"
	//g := `Miss Moneypenny says, "Helloooooo, James."`
	//h := `Q says,
	//"I have some new toys for you, James."
	// `
	fmt.Printf("%v\n", a3)
	fmt.Printf("%v\n", b3)

	// Sprint

	const name, age = "Kim", 22
	s := fmt.Sprintf("%s tiene %d years.\n", name, age)

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
