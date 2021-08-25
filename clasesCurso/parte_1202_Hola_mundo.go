package clasesCurso

import "fmt"

func Clase1202() {
	/*
			Hola mundo

	●	Estructura básica de un programa en Go:
		○	package main
		○	func main
	■	Punto de entrada o inicio a tu programa
	■	Cuando tu código sale de la función main, tu programa ha terminado
		●	Cantidad variable de parámetros
			○	el “...<algún tipo>” es como especificamos un variadic parameter
			○	El tipo “interface{}” es una interfaz vacía
	■	Todos los tipos son de tipo “interface{}”
		○	Entonces el parámetro  “...interface{}” significa “ingresa la cantidad de argumentos que quieras”
			●	Deshaciéndose de los returns
		○	uso del caracter “_” para deshacerte de los returns
			●	No puedes tener variables sin usar en tu código
		○	Esto es contaminación de código
		○	El compilador no lo permite
			●	Usamos esta notación en Go
		○	paquete.Identificador
	■	ejemplo: fmt.Println()
			●	Leeremos: “del paquete fmt Estoy usando la función Println”
				○	Un identificador es el nombre de una variable, constante, función
			●	paquetes
				○	Código que está ya escrito el cual puedes usar
				○	imports

	*/
	fmt.Println("Clase1202")
}
