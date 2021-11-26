package clasesCurso

import (
	"fmt"
)

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, `, dice "Hola."`)
	q := true
	return p, q

}

//************-----------------------------**************************************************
func Clase4201() {

	multi_line := `Funciones
	Sintaxis
	func (receptor) identificador(parámetros) (returns) { código }
	Conoce la diferencia entre parámetros y argumentos
		●	Definimos las funciones con parámetros (si lleva alguno)
		●	Llamamos las funciones y les pasamos argumentos (si lleva alguno)
	Todo en Go es PASADO POR VALOR
	Propósito de las funciones
		●	Abstraer código
		●	Reutilización del código 
	
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := saludar("Eduar", "Tua")
	fmt.Println(x)
	fmt.Println(y)

}

//************-----------------------------**************************************************
func Clase4202() {

	multi_line := `Parámetros variados
	Puedes crear una func el cual toma un número ilimitado de argumentos. Cuando haces esto, 
	se conoce como “parámetro variado” o en inglés como “variadic parameter.” Cuando usas el 
	operador que forma parte de los elementos de léxico “...T” para indicar un parámetro variado 
	(donde “T” representa algún tipo).
			`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------
	x := sum4202(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("El total almacenado es", x)
}

func sum4202(x ...int) int { // ... -> cantidad variable de parametros
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el índice", i, "suma", v, "al total, quedando", suma)
	}
	fmt.Println("El total es", suma)
	return suma
}

//func (r receptor) identificador(parámetros) (retorno(s)) { Código }

//************-----------------------------**************************************************
func Clase4203() {

	multi_line := `Desplegando un slice
	Cuando tienes un slice de algún tipo, puedes pasar cada uno de los valores 
	individuales en un slice usando el operador “...”
	
	
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------
	x := sum1("james")
	fmt.Println("The total is", x)
}

func sum1(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}

//************-----------------------------**************************************************
func Clase4204() {

	multi_line := `Defer
	Una declaración "defer" (diferir) invoca a una función cuya ejecución es diferida para 
	el momento en el que la función donde está contenida retorna, en cualquiera de los siguientes 
	casos: o la función ejecuta una  declaración de retorno, o llega al final de su cuerpo o porque 
	la gorutina (goroutine) correspondiente está en pánico (panicking).

	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------
	defer foo4202()
	bar4202()
}

func foo4202() {
	fmt.Println("foo4202")
}

func bar4202() {
	fmt.Println("bar4202")
}

//************-------------Métodos----------------**************************************************
type persona4205 struct {
	nombre   string
	apellido string
}

type agenteSecreto4205 struct {
	persona4205
	lpm bool
}

// func (r receptor) identificador(parámetros) (retorno(s)) { código }

func (s agenteSecreto4205) presentar() {
	fmt.Println("Hola, soy", s.nombre, s.apellido)
}

func Clase4205() {

	multi_line := `Métodos
	Un método no es más que una FUNC adjuntada a un TIPO. Cuando adjuntas una función 
	a un tipo es un método de ese tipo. Adjuntas una función a un tipo con el RECEPTOR.
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------
	sa1 := agenteSecreto4205{
		persona4205: persona4205{
			"Eduar",
			"Tua",
		},
		lpm: true,
	}

	sa2 := agenteSecreto4205{
		persona4205: persona4205{
			"Miss",
			"Moneypenny",
		},
		lpm: true,
	}

	fmt.Println(sa1)
	sa1.presentar()
	sa2.presentar()
}

//************----------Interfaces y polimorfismo-------------------**************************************************
type person4206 struct {
	first string
	last  string
}

type secretAgent4206 struct {
	person4206
	ltk bool
}

func (s secretAgent4206) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person4206) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human4206 interface {
	speak()
}

func bar4206(h human4206) {
	switch h.(type) {
	case person4206:
		fmt.Println("I was passed into barrrrrr", h.(person4206).first)
	case secretAgent4206:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent4206).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog4206 int

func Clase4206() {

	multi_line := `Interfaces y polimorfismo
	En Go, los valores pueden ser de más de un tipo. Una interfaz permite a un valor 
	ser de más de un tipo. Creamos una interfaz usando la sintaxis : “palabra clave 
	identificador tipo” entonces para una interfaz sería: “type humano interface” 
	Luego definimos cuáles método(s) debe tener un tipo para implementar esa interfaz. 
	Si un TIPO tiene los métodos requeridos, el cual podrían ser ninguno, 
	(la interfaz vacía denotada por interface{}), entonces ese TIPO implícitamente 
	implementa la interfaz y es también de ese tipo de interfaz. En Go, los valores 
	pueden ser más de un tipo.
		
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------
	sa1 := secretAgent4206{
		person4206: person4206{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent4206{
		person4206: person4206{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person4206{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar4206(sa1)
	bar4206(sa2)
	bar4206(p1)

	// conversion
	var x hotdog4206 = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

//************--------------Funciones Anónimas---------------**************************************************
func Clase4207() {

	multi_line := `Funciones Anónimas
	
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------
	foo4207()

	func() {
		fmt.Println("La función anónima se ejecutó")
	}()

	func(x int) {
		fmt.Println("El significado de la vida es:", x)
	}(42)
}

func foo4207() {
	fmt.Println("foo se ejecutó")
}

//************------------Expresión función-----------------**************************************************
func Clase4208() {

	multi_line := `
	
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	f := func() {
		fmt.Println("Mi primera expresión función")
	}
	f()

	g := func(x int) {
		fmt.Println("El año del descubrimiento de América fue:", x)
	}
	g(1492)
}

//************-----------Retornando una func------------------**************************************************
func Clase4209() {

	multi_line := `Retornando una func
	
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	fmt.Println(bar4209()())

}

func bar4209() func() int {
	return func() int {
		return 451
	}
}

//************------------Callback-----------------**************************************************
func Clase4210() {

	multi_line := `●	Pasando una func como un argumento
	●	La programación funcional no es algo que se recomienda hacer con Go, sin embargo, 
		es bueno estar al tanto de los callbacks.
	●	idiomatic go: escribe código que sea claro, simple y legible.
	
	
	`
	fmt.Printf("%s", multi_line)
	//operadores matematicos----------------------------------------------------------------------
	x := 0
	fmt.Println(x)

	x++
	fmt.Println(x)

	x += 42
	fmt.Println(x)

	x -= 3
	fmt.Println(x)

	x--
	fmt.Println(x)

	//SOLO la funcion suma----------------------------------------------------------------------
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum4210(ii...)
	fmt.Println("all numbers", s)
	//Numeros pares----------------------------------------------------------------------
	s2 := even4210(sum4210, ii...)
	fmt.Println("even numbers", s2)
	//Numeros impares----------------------------------------------------------------------
	s3 := odd4210(sum4210, ii...)
	fmt.Println("odd numbers", s3)
}

func sum4210(xi ...int) int {
	// fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even4210(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd4210(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

//************--------------Closure---------------**************************************************
func Clase4211() {

	multi_line := `●	En un scope que encierra otros scopes
	○	Las variables declaradas en el scope externo son accesibles en los scopes internos.
	●	Los closures nos ayudan a limitar el scope de las variables
	
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

//************--------------Recursividad---------------**************************************************
func Clase4212() {

	multi_line := `Recursividad
	●	Una func que se llama a ella misma
	●	Ejemplo factorial

	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)

	n2 := loopFact(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

//************-----------------------------**************************************************
func Clase4213() {

	multi_line := `
	
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

}
