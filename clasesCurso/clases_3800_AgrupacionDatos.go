package clasesCurso

import (
	"fmt"
)

func Clase3800() {

	/*
		●	Arreglo (Array)
			○	Es una secuencia enumerada de elementos de un mismo tipo.
			○	No cambia en tamaño
			○	Usado por funcionalidades internas de Go; generalmente no es recomendado que lo uses en tu código.

		●	Slice
			○	Están construidos sobre el tipo arreglo.
			○	Los valores que están contenidos en un slice son del mismo tipo.
			○	Este si puede cambiar en tamaño.
			○	Tiene una longitud y una capacidad.

		●	map
			○	Almacenamiento key / value
			○	Es un grupo de elementos de un tipo, llamado element type, sin orden de agrupación, indexados por un conjunto único de keys de otro tipo, llamado key type.

		●	struct
			○	Una estructura de datos
			○	Un tipo compuesto
			○	Nos permite poner juntos valores de diferentes tipos.


	*/

	multi_line := `	
●	Arreglo (Array)
	○	Es una secuencia enumerada de elementos de un mismo tipo.
	○	No cambia en tamaño
	○	Usado por funcionalidades internas de Go; generalmente no es recomendado que lo uses en tu código.

●	Slice
	○	Están construidos sobre el tipo arreglo.
	○	Los valores que están contenidos en un slice son del mismo tipo.
	○	Este si puede cambiar en tamaño.
	○	Tiene una longitud y una capacidad.

●	map
	○	Almacenamiento key / value
	○	Es un grupo de elementos de un tipo, llamado element type, sin orden de agrupación, indexados por un conjunto único de keys de otro tipo, llamado key type.

●	struct
	○	Una estructura de datos
	○	Un tipo compuesto
	○	Nos permite poner juntos valores de diferentes tipos.
`
	fmt.Printf("%s", multi_line)

}

func Clase3801() {

	/*
		Arreglos (arrays)
		Los arreglos son mayormente usados como un bloque constructor en el lenguaje de programación Go.
		En algunas circunstancias, podríamos usar un array, pero en la mayoría de los casos la
		recomendación es usar slices en vez de arreglos.


	*/

	multi_line := ` Arreglos (arrays)
	Los arreglos son mayormente usados como un bloque constructor en el lenguaje de programación Go. 
	En algunas circunstancias, podríamos usar un array, pero en la mayoría de los casos la 
	recomendación es usar slices en vez de arreglos.

`
	fmt.Printf("%s", multi_line)

	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	fmt.Println(len(x))

}

func Clase3802() {

	/*
				Slice - literal compuesto
		UN SLICE almacena VALORES del mismo TIPO. Si quisiéramos almacenar todos nuestros números favoritos
		usamos un SLICE de ints. Si quisiera almacenar todas mis comidas favoritas usaría un SLICE de strings.
		Usaremos un LITERAL COMPUESTO para crear un slice. Un literal compuesto es cuando colocamos el TIPO
		seguido de LLAVES y luego colocamos los valores en el área dentro de las llaves.

	*/

	multi_line := `Slice - literal compuesto
	UN SLICE almacena VALORES del mismo TIPO. Si quisiéramos almacenar todos nuestros números favoritos 
	usamos un SLICE de ints. Si quisiera almacenar todas mis comidas favoritas usaría un SLICE de strings.
	Usaremos un LITERAL COMPUESTO para crear un slice. Un literal compuesto es cuando colocamos el TIPO 
	seguido de LLAVES y luego colocamos los valores en el área dentro de las llaves.
`
	fmt.Printf("%s", multi_line)

	//tipo{elementos} //COMPOSITE LITERAL
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		//fmt.Println(i, " ", v)
		fmt.Printf("%v %T\n", i, v)

	}

}

func Clase3803() {

	/*


	 */

	multi_line := `Slice - for range
	Podemos iterar sobre los valores en un slice con la cláusula range. También le podemos agregar 
	elementos a un slice mediante el uso del índice.	
`
	fmt.Printf("%s", multi_line)

	//tipo{elementos} //COMPOSITE LITERAL
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(cap(x)) // capacidad del slice

	for i, v := range x {
		//fmt.Println(i, " ", v)
		fmt.Printf("%v %T\n", i, v)

	}

}

func Clase3804() {

	/*


	 */

	multi_line := `Slice - Dividiendo un slice
	Podemos dividir un Slice, lo que quiere decir que podemos cortar y desechar partes de un slice. 
	Hacemos esto con el operador dos puntos ( : ).
		
`
	fmt.Printf("%s", multi_line)

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// dividir o seleccionar solo un rango
	fmt.Println("dividir o seleccionar solo un rango")
	fmt.Println(x[:]) // slice completo
	fmt.Println(x[1:3])
	fmt.Println(x[2:4])

	for i, v := range x {
		fmt.Println(i, " ", v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}

func Clase3805() {

	/*


	 */

	multi_line := `Slice - Añadiendo a datos
	Para añadir valores a un slice, debemos usar la función especial integrada append. Esta función retorna un slice del mismo tipo. 
	
		
`
	fmt.Printf("%s", multi_line)

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 44, 55, 66)
	fmt.Println(x)

	y := []int{333, 444, 666}
	x = append(x, y...)
	fmt.Println(x)

}

func Clase3806() {

	/*


	 */

	multi_line := `Slice - Eliminado elementos de un slice
	Podemos eliminar elementos de un slice usando append y slicing (dividiendo). 
	Este es un maravilloso y elegante ejemplo de porqué Go es súper cool y cómo 
	provee facilidad de programación.`
	fmt.Printf("%s", multi_line)

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 44, 55, 66)
	fmt.Println(x)

	y := []int{333, 444, 666}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

}

func Clase3807() {

	/*


	 */

	multi_line := `Slice - make
	Los slices son construidos sobre los arreglos. Un slice es dinámico, así este crecerá en tamaño. 
	Sin embargo, el arreglo subyacente, no crece en tamaño. Cuando creamos un slice, podemos usar la 
	función predefinida interna make para especificar que tan grande debería ser nuestro slice y también 
	qué tan grande debería ser el arreglo subyacente. Esto puede mejorar un poco el desempeño del programa.
	○	make([]T, length, capacity) 
	○	make([]int, 50, 100) 
	`
	fmt.Printf("%s", multi_line)

	x := make([]int, 10, 12) // 10 igual a posiciones ,12 = capacidad
	fmt.Println(x)
	fmt.Println(len(x)) // 10
	fmt.Println(cap(x)) // 100
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423) // agrega el valor en posicion 11 que se activa hasta esta asignacion

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425) // si la capacidad es superada, la capacidad se duplica automaticamente

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func Clase3808() {

	/*


	 */

	multi_line := `Slice - slices multidimensionales
	Un slice multidimensional es como una hoja de cálculo. Puede tener un slice de slices de algún tipo. 
	Suena confuso?.
	
	
	`
	fmt.Printf("%s", multi_line)

	et := []string{"Eduar", "Tua", "Crossfit", "Baseball", "Montañismo"}
	fmt.Println(et)
	jl := []string{"Jacinto", "Lopez", "Correr", "Nadar", "Bailar"}
	fmt.Println(jl)

	vp := [][]string{et, jl}
	fmt.Println(vp)
}

func Clase3809() {

	/*


	 */

	multi_line := `Slice - El arreglo subyacente
	Subyacente a cada slice habrá un arreglo (array). Un slice es realmente una estructura de datos el cual tiene tres partes: 
	1.	Un puntero a un arreglo
	2.	Longitud (len)
	3.	Capacidad (cap)
	`
	fmt.Printf("%s", multi_line)

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	y := append(x[:2], x[3:]...) // se utiliza el mismo arreglo o puntero ya que el slice es menor
	fmt.Println(x)               // imprime slice x modificado por y
	fmt.Println(y)
}

func Clase3810() {

	multi_line := `Mapa (map) - introducción
	Un mapa es un tipo de dato de almacenamiento llave-valor. Esto quiere decir que almacenas algún valor y 
	accedes al mismo con una llave. Por ejemplo, podría almacenar el valor “numeroTel” y acceder a él con la 
	llave “nombreAmigo”. De esta manera puedo ingresar el nombre de mi amigo y el mapa me retornará su número 
	telefónico. La sintaxis para el mapa es map[llave]valor. La llave puede ser de cualquier tipo que permita 
	comparación (por ejemplo, podría usar un string o un int, ya que puedo comparar si dos strings son iguales 
	o si dos enteros son iguales. Es importante denotar que los mapas son desordenados. Si imprimes todas las 
	llaves y valores de un mapa se hará de forma aleatoria. El idioma comma ok es también cubierto en este 
	video. Un mapa es la estructura de datos perfecta cuando necesitas buscar datos de manera rápida. 
	
	`
	fmt.Printf("%s", multi_line)

	m := map[string]int{ // mapa ejemplo
		"Batman": 32, // tipo : valor
		"Robin":  27,
	}
	fmt.Println(m)

	fmt.Println(m["Batman"]) //acceso a llave especifica de m

	fmt.Println(m["Eduar"]) //acceso a llave que no existe

	v, ok := m["Eduar"] // idioma, ok // verdadero si existe y falso si no.
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Eduar"]; ok {
		fmt.Println("Impresión desde el if", v)
	}
}

func Clase3811() {

	multi_line := `Map - Agregar elemento y range
	Aquí mostramos como agregar un elemento a un mapa. También muestro cómo usar el 
	ciclo for range para imprimir las llaves y valores de un mapa.
	`
	fmt.Printf("%s", multi_line)

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 33 // agregar llave y valor a m

	if v, ok := m["Barnabas"]; ok { // no entra al ciclo no imprime nada
		fmt.Println(v)
	}

	for k, v := range m { // recorre el map m
		fmt.Println(k, v) // imprime llave - valor
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}

func Clase3812() {

	multi_line := `Map - Agregar elemento y range
	Aquí mostramos como agregar un elemento a un mapa. También muestro cómo usar 
	el ciclo for range para imprimir las llaves y valores de un mapa.
	
	`
	fmt.Printf("%s", multi_line)

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James") // borra James y su valor de m
	fmt.Println(m)

	if v, ok := m["Robin"]; ok {
		fmt.Println("Se borró la llave con valor", v)
		delete(m, "Robin")
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 33

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
