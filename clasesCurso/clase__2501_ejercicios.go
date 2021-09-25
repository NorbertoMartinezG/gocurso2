package clasesCurso

import (
	"fmt"
)

func Clase2501() {
	/*
	   Ejercicio práctico #1
	   	Escribe un programa que imprima un número en decimal, binario, y hexadecimal

	*/
	fmt.Println("Escribe un programa que imprima un número en decimal, binario, y hexadecimal")
	a := 42
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}

func Clase2502() {
	/*
		Ejercicio práctico #2
			Usando los siguientes operadores, escribe expresiones y asigna sus valores a variables:
			a.	==
			b.	<=
			c.	>=
			d.	!=
			e.	<
			f.	>
			Imprime los valores de las variables.

	*/
	fmt.Println("Usando los siguientes operadores, escribe expresiones y asigna sus valores a variables:")
	fmt.Println("a.	==,	b.	<=,	c.	>=,	d.	!=,	e.	<,	f.	>")

	fmt.Println("a := (42 == 42)")
	a := (42 == 42)
	fmt.Println(a)
	fmt.Println("b := (42 <= 43)")
	b := (42 <= 43)
	fmt.Println(b)
	fmt.Println("c := (42 >= 43)")
	c := (42 >= 43)
	fmt.Println(c)
	fmt.Println("d := (42 != 43)")
	d := (42 != 43)
	fmt.Println(d)
	fmt.Println("e := (42 < 43)")
	e := (42 < 43)
	fmt.Println(e)
	fmt.Println("f := (42 > 43)")
	f := (42 > 43)
	fmt.Println(f)

}

func Clase2503() {
	/*
		Crea constantes CON TIPO y SIN TIPO. Imprime el valor de las mismas.
	*/
	fmt.Println("Crea constantes CON TIPO y SIN TIPO. Imprime el valor de las mismas.")

	const (
		a     = 42
		b int = 43
	)
	fmt.Println(a, b)

}

func Clase2504() {
	/*
		Escribe un programa que
		●	Asignar un int a una variable
		●	Imprímelo en decimal, binario, y hex
		●	Has shifts de bits de ese int una posición a la izquierda y asigna eso a una variable
		●	Imprime esa variable en decimal, binario, y hex

	*/
	multi_line := `Escribe un programa que
	●	Asignar un int a una variable
	●	Imprímelo en decimal, binario, y hex
	●	Has shifts de bits de ese int una posición a la izquierda y asigna eso a una variable
	●	Imprime esa variable en decimal, binario, y hex
`
	fmt.Printf("%s", multi_line)

	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Println("shifts de bits de 42 int una posición a la izquierda y asigna eso a una variable")
	fmt.Printf("%d\t%b\t%#x", b, b, b)

}

func Clase2505() {
	/*
		Crea una variable de tipo string usando un string literal no interpretado (raw string literal). Imprímelo.
	*/
	multi_line := `Crea una variable de tipo string usando un string literal no interpretado (raw string literal). Imprímelo.`
	fmt.Printf("%s", multi_line)

	a := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`
	fmt.Println(a)

}

const (
	a00 = iota
	a11 = 2017 + iota
	b11 = 2017 + iota
	c11 = 2017 + iota
	d11 = 2017 + iota
)

func Clase2506() {
	/*
		Crea una variable de tipo string usando un string literal no interpretado (raw string literal). Imprímelo.
	*/
	multi_line := `Usando iota, crea 4 constantes para los PRÓXIMOS 4 años. Imprime los valores de las constantes.`
	fmt.Printf("%s", multi_line)

	fmt.Println("a00")
	fmt.Println(a00)
	fmt.Println(a11)
	fmt.Println(b11)
	fmt.Println(c11)
	fmt.Println(d11)

}
