package clasesCurso

import (
	"fmt"
)

func Clase2601() {

	fmt.Println("Entendiendo el control de flujo")

	/*
			La computadora lee los programas de cierta forma, así como nosotros leemos libros de cierta forma.
			Cuando, como humanos, leemos libros, en culturas occidentales, lo hacemos de la parte frontal hacia la trasera del libro,
			de izquierda a derecha y de arriba hacia abajo. Cuando las computadoras leen el código de un programa, lo hacen de arriba
			hacia abajo. Esto es conocido como lectura en SECUENCIA y a su vez conocido como control de flujo secuencial. Hay otras dos
			declaraciones el cuál pueden afectar cómo la computadora lee el código. Una computadora puede encontrarse con un CICLO o LOOP.
			Si se encuentra con uno de esos entrará en un bucle e iterará sobre él de una manera específica. Por eso es también conocido
			como control de flujo ITERATIVO. Finalmente, hay también control de flujo CONDICIONAL. Cuando la computadora se encuentra
			con una CONDICIÓN, como una "declaración if" o una "declaración de switch" la computadora tomará acciones diferentes
			dependiendo de alguna condición. Entonces las tres formas en la que la computadora lee el código son: SECUENCIAL, LOOP, CONDICIONAL

		●	secuencia
		●	loop / iterativo
			○	for loop
				■	init, cond, post
			○	bool (con while)
				■	infinite
			○	Con do-while
				■	break
			○	contínuo
			○	anidado
		●	condicionales
			○	Declaraciones switch / case / default
				■	Sin caída predeterminada
				■	Creando caída (creating fall-through)
				■	Casos múltiples
				■	Casos pueden ser expresiones
		●	Evaluar a true, ellos corren
				■	tipo
			○	if
				■	El operador de negación
					●	!
				■	Declaración de inicialización
				■	if / else
				■	if / else if / else
				■	if / else if / else if / … / else


	*/

}

func Clase2602() {

	fmt.Println("Ciclo - init, condition, post")

	/*
			Ciclo - init, condition, post
		●	Ciclo for
		○	inicialización, condition, post

	*/

	// for init; condition; post {}
	fmt.Println("Ciclo for de 0 a 10")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)

	}

}

func Clase2603() {

	fmt.Println("Ciclos anidados")

	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("El cilo externo: %d\t El ciclo interno: %d\n", i, j)

		}
	}

	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("El ciclo externo: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t El ciclo interno: %d\n", j)

		}
	}

}

func Clase2604() {

	fmt.Println("Ciclo - Declaración for")

	/*
		Hay tres formas con las que puedes construir ciclos for en Go - todas usan la palabra clave “for”:
		●	for init; condición; post { }
		●	for condición { }
		●	for { }
		Respecto a la documentación de la declaración “for”
		●	Especificaciones del lenguaje
		●	effective go


	*/
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("Listo.")

}

func Clase2605() {

	fmt.Println("Ciclo - break & continue")

	/*
		Ciclo For
		●	break
		●	continue
		Encontrando el resto, también conocido como modulus
		●	%


	*/

	x1 := 83 / 40
	y := 83 % 40
	fmt.Println(x1, y)

	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)

	}
	fmt.Println("done.")

}

func Clase2606() {

	fmt.Println("Ciclo - Imprimiendo ascii")

	/*
		Podemos usar impresión con formato para imprimir:
		●	Valor decimal con %d
		●	Valor hexadecimal con %#x
		●	Código unipoint %#U
		●	Una tabulación con  \t
		●	Una línea nueva con \n


	*/
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}

}

func Clase2607() {

	fmt.Println("Condicional - if, else if, else")

	x := 434
	if x == 40 {
		fmt.Println("El valor es 40")
	} else if x == 41 {
		fmt.Println("El valor es 41")
	} else if x == 42 {
		fmt.Println("El valor es 42")
	} else if x == 43 {
		fmt.Println("El valor es 43")
	} else {
		fmt.Println("El valor no es 40, 41, 42, ni 43")
	}

}

func Clase2610() {

	fmt.Println("Condicional - Declaración switch")

	/*
		Ciclo, condicional, módulo
		Iteramos para mejorar nuestras vidas. Esto se cumple para cualquier cosa que estemos haciendo.
		Escribimos código con errores antes de que escribamos código sin errores. Un profesor una vez llamó a esto
		“perfección de la imperfección.”  “Tú eres perfecto de la manera que eres, y siempre habrá oportunidad de mejora.”
		El punto es-  el código que escribimos para imprimir número pares puede se hecho de una mejor manera. Vamos a escribir
		código que es más legible. Esto servirá como una revisión de los conceptos que estamos aprendiendo: loops, condicionales,
		sentencia if y el operador módulo.

				Declaración Switch
		●	Declaraciones switch / case / default
		○	Sin caer en la declaración fall-through
		○	Creando  fall-through
		○	multiples casos
		○	cases pueden ser expresiones
		■	Evalúan a verdadero, corren.



	*/
	n := "Pera"
	switch n {
	case "Manzana", "Pera", "Mango":
		fmt.Println("Varias frutas")
		fallthrough // imprime el siguiente caso aunque sea falso
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("Esta es la q")
	default:
		fmt.Println("Esta es default")
	}

}

func Clase2611() {

	fmt.Println("Condicional - documentación de switch")

	/*
		Condicional - documentación de switch
		Sentirse cómodo mientras leemos la documentación es importante. Dedicar un poco de tiempo para que veamos la documentación de Go es importante, te será útil ver cómo se lee e interpreta la misma y te hará sentir más cómodo con ella.
		video: 047
		Operadores lógicos condicionales
		A qué evalúan las siguientes expresiones :
		●	fmt.Println(2 == 2 && 3 == 3)
		●	fmt.Println(true && false)
		●	fmt.Println(5 == 5 || 6 == 6)
		●	fmt.Println(true || false)
		●	fmt.Println(!true)


	*/
	//fmt.Printf("true && true\t %v\n", true && true)
	fmt.Printf("true && false\t %v\n", true && false)
	//fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}
