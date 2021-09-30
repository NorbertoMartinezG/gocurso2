package clasesCurso

import (
	"fmt"
)

func Clase2701() {

	fmt.Println("Imprime todos los números del 1 al 10,000")

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func Clase2702() {

	multi_line := `Ejercicio práctico #2
	Imprime el rune code point de las letras del alfabeto en mayúsculas tres veces. 
	La salida debe ser como esto:
	65
		U+0041 A
		U+0041 A
		U+0041 A
	66
		U+0042 B
		U+0042 B
		U+0042 B 
	 … hasta el resto de letras en el alfabeto.
`
	fmt.Printf("%s", multi_line)

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func Clase2703() {

	multi_line := `Ejercicio Práctico #3
	Crea un ciclo usando la siguiente sintaxis
	●	for condición { }
	Haz que imprima los años que has vivido.
	
`
	fmt.Printf("%s", multi_line)

	bd := 1991
	for bd <= 2021 {
		fmt.Println(bd)
		bd++
	}
}

func Clase2704() {

	multi_line := `Ejercicio Práctico #4
	Crea un ciclo for usando esta sintaxis
	●	for { }
	Haz que imprima los años que has vivido.
	
`
	fmt.Printf("%s", multi_line)

	bd := 1991
	for {
		if bd > 2021 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}

func Clase2705() {

	multi_line := `Ejercicio Práctico #5
	Imprime el resto o módulo, el cual es resultado de dividir entre 4 cada número en el rango de 10 y 100.	
`
	fmt.Printf("%s", multi_line)

	for i := 10; i <= 100; i++ {
		fmt.Printf("Cuando dividimos %v entre 4, el resto (también módulo) es %v\n", i, i%4)
	}
}

func Clase2706() {

	multi_line := `Crea un programa que muestre el “if statement” en acción
	x := "Batman"
	if x == "Batman" {
		fmt.Println(x)
	}
`
	fmt.Printf("%s", multi_line)

	x := "Batman"
	if x == "Batman" {
		fmt.Println(x)
	}
}

func Clase2707() {

	multi_line := `Usando el ejercicio anterior, crea un programa que use “else if” y “else”.
	x := "Rsdfg"
	if x == "Robin" {
		fmt.Println(x)
	} else if x == "Batman" {
		fmt.Println("batbatbat", x)
	} else {
		fmt.Println("Ningún súper héroe")
	}
`
	fmt.Printf("%s", multi_line)

	x := "Rsdfg"
	if x == "Robin" {
		fmt.Println(x)
	} else if x == "Batman" {
		fmt.Println("batbatbat", x)
	} else {
		fmt.Println("Ningún súper héroe")
	}
}

func Clase2708() {

	multi_line := `Crea un programa que use una declaración switch sin expresión especificada.
	switch {
	case false:
		fmt.Println("No debería imprimir")
	case true:
		fmt.Println("Debería imprimir")
	}
`
	fmt.Printf("%s", multi_line)

	switch {
	case false:
		fmt.Println("No debería imprimir")
	case true:
		fmt.Println("Debería imprimir")
	}
}

func Clase2709() {

	multi_line := `Crea un programa que use una declaración switch con la expresión de switch
	especificada como una variable de TIPO string y el IDENTIFICADOR “deporteFav”.
	
	deporteFav := "béisbol"
	switch deporteFav {
	case "béisbol":
		fmt.Println("Ve al estadio")
	case "natación":
		fmt.Println("Ve a la piscina")
	case "crossfit":
		fmt.Println("Te quiero ver en los crossfit games.")
	}
`
	fmt.Printf("%s", multi_line)

	deporteFav := "béisbol"
	switch deporteFav {
	case "béisbol":
		fmt.Println("Ve al estadio")
	case "natación":
		fmt.Println("Ve a la piscina")
	case "crossfit":
		fmt.Println("Te quiero ver en los crossfit games.")
	}
}
