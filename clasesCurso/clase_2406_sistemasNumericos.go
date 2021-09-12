package clasesCurso

import "fmt"

func Clase2406() {
	fmt.Println("Clase 2406 sistemas numericos")

	/*
		Sistemas numéricos
			Como humanos, tenemos diferentes sistemas para expresar las cantidades de algo. Usando el sistema numérico decimal,
			podemos decir que tenemos 7 naranjas; 42 manzanas o 20 dólares. Otros sistemas numéricos usados en computación incluyen el
			sistema numérico binario y el sistema numérico hexadecimal.

			Constantes
				●	Un simple valor que no cambia.
				●	Sólo existe en el momento de compilación..
				●	Hay constantes con TIPO y SIN TIPO
					○	const hola = "Hola, mundo"
					○	const typedHello string = "Hello, World"
			●	Constante SIN TIPO
				○	 Un valor constante que no tiene un tipo fijo
					■	“constante de algún tipo”
					■	No es forzada a obedecer las reglas estrictas que previenen combinar diferentemente valores con un tipo.
				○	Una constante sin tipo puede ser implícitamente convertida por el compilador.
			●	Es este concepto de constante sin tipo lo que hace posible que usemos constantes en Go con libertad.
				○	Muy útil, por ejemplo
					■	Cuál es el tipo de 42?
						●	int?
						●	uint?
						●	float64?
			■	Si no tuviéramos constante SIN TIPO (constantes de algún tipo), tendríamos que hacer conversión en cada valor literal que usamos.
			●	Y eso no sería muy agradable

	*/

}

func Clase2407() {
	fmt.Println("Clase 2407 constantes")
	const a = 42
	const b = 42.32
	const c = "Eduar Tua"

	type nombre string // definiendo variable tipo nombre tipo String

	var d nombre

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a) //imprime el tipo asignado por el sistema
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	d = c // error d es tipo nombre y c tipo string -- const c string = "Eduar Tua"
	fmt.Println(d)

	const x = 42
	const y float64 = 43.2
	type hotdog int
	type hotcat float64

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf("%T\n", hotdog(x))
	//	fmt.Printf("%T\n", hotdog(y))

	fmt.Printf("%T\n", hotcat(x))
	fmt.Printf("%T\n", hotcat(y))
}

func Clase2408() {

	/*
		Iota
		Dentro de una declaración, el identificador pre-declarado iota representa una sucesión de
		constantes enteras sin tipo. Es reiniciado a 0 cada vez que la palabra constante aparece en el
		código. Puede ser usada para construir un conjunto de constantes relacionadas:

	*/
	const (
		a = iota // inicia un int en 0
		b        // continua con la numeracion dentro de estos const hasta que aparece la palabra const
		c
	)

	const ( //aqui se reinicia la numeracion
		d = iota
		e
		f
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func Clase2409() {
	/*
		Bit shifting
		Bit shifting es cuando desplazan dígitos binarios a la izquierda o a la derecha.
		Podemos usar bit shifting junto con iota, para construir constantes creativas.

	*/
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x) // imprimiendo en binario y decimal

	y := x << 1 // << >> operadores bit shifting en este caso muestra "1000" que en binario representa a 8 ya que es el que sigue de 4 o x en esa variable
	fmt.Printf("%d\t\t%b", y, y)

	fmt.Println("")
	kb := 1024      //kilobyte
	mb := 1024 * kb //megabyte
	gb := 1024 * mb //gibyte

	fmt.Printf("%d\t\t\t%b\n", kb, kb) // imprime decial, y binario
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	const (
		_   = iota             // aqui se desecha iota con _ para reinciar
		kb1 = 1 << (iota * 10) //iota = 1
		mb1 = 1 << (iota * 10) //iota = 2
		gb1 = 1 << (iota * 10) //iota = 3
	)

	fmt.Printf("%d\t\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t\t%b\n", mb1, mb1)
	fmt.Printf("%d\t\t%b\n", gb1, gb1)

}
