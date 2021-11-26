package clasesCurso

import (
	"fmt"
)

func Clase3901() {

	multi_line := `Ejercicio Práctico #1
	●	Usando un COMPOSITE LITERAL: 
	●	Crea un ARREGLO el cual tenga 5 VALORES de TIPO int
	●	Asigna VALORES a cada posición dada por los índices. 
	●	Itera con Range sobre el arreglo e imprime los valores. 
	●	Usando el paquete fmt
	○	Imprime el TIPO del arreglo
		`
	fmt.Printf("%s", multi_line)

	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func Clase3902() {

	multi_line := `Ejercicio Práctico #2
	●	Usando un COMPOSITE LITERAL: 
	●	Crea un SLICE de TIPO int
	●	Asigna 10 VALORES 
	●	Haz Range sobre el slice e imprime los valores.
	●	Usando format para imprimir
	○	Imprime el TIPO del slice
		`
	fmt.Printf("%s", multi_line)

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func Clase3903() {

	multi_line := `Ejercicio Práctico #3
	Usando el código del ejercicio anterior, usa SLICING para crear los siguientes nuevos slices el cual serán impresos:
	●	[42 43 44 45 46]
	●	[47 48 49 50 51]
	●	[44 45 46 47 48]
	●	[43 44 45 46 47]
		`
	fmt.Printf("%s", multi_line)

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}

func Clase3904() {

	multi_line := `Ejercicio Práctico #4
	Sigue los siguientes pasos:
	●	Comienza con este slice
	○	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	●	Agrégale el siguiente valor
	○	52
	●	Imprime el slice
	●	En UNA DECLARACIÓN agrega al slice los siguientes valores
	○	53
	○	54
	○	55
	●	Imprime el slice
	●	Agregale al Slice los siguientes valores
	○	y := []int{56, 57, 58, 59, 60}
	●	print out the slice
		`
	fmt.Printf("%s", multi_line)

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func Clase3905() {

	multi_line := `Ejercicio Práctico #5
	Para BORRAR de un slice, usamos APPEND en conjunto con SLICING(dividir). Para este ejercicio sigue los siguientes pasos:
	●	Comienza con un slice
	○	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	●	Usa APPEND & SLICING para obtener los siguientes valores el cual se los debes asignar a una variable “y” y luego imprimir:
	○	[42, 43, 44, 48, 49, 50, 51]
		`
	fmt.Printf("%s", multi_line)

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x[:3], x[6:]...) // [42, 43, 44, 48, 49, 50, 51]
	fmt.Println(x)
}

func Clase3906() {

	multi_line := `Ejercicio Práctico #6
	Crea un slice para almacenar los nombres de todos los estados en los Estados Unidos de América. 
	¿Cuál es la longitud del slice? ¿Cuál es la capacidad del Slice? Imprime todos los valores con su  
	índice de posición, sin utilizar la cláusula range. Aquí la lista de los estados:
	Alabama,  Alaska,  Arizona,  Arkansas,  California,  Colorado,  Connecticut,  Delaware,  
	Florida,  Georgia,  Hawaii,  Idaho,  Illinois,  Indiana,  Iowa,  Kansas,  Kentucky, 
	Louisiana,  Maine,  Maryland,  Massachusetts,  Michigan,  Minnesota,  Mississippi, 
	Missouri,  Montana,  Nebraska,  Nevada, New Hampshire,  New Jersey,  New Mexico, 
	New York,  North Carolina,  North Dakota,  Ohio,  Oklahoma,  Oregon,  Pennsylvania,
	Rhode Island,  South Carolina,  South Dakota,  Tennessee,  Texas,  Utah,  Vermont, 
	Virginia,  Washington,  West Virginia,  Wisconsin,  Wyoming, 
	`
	fmt.Printf("%s", multi_line)

	x := make([]string, 50, 50)
	x = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func Clase3907() {

	multi_line := `Ejercicio Práctico #7
	Crear un slice de slice de string ([][]string). Almacena los siguientes valores en un slice multi-dimensional:
	
	"Batman", "Jefe", "Vestido de negro"
	"Robin", "Ayudante", "Vestido de colores"
	
	Haz range sobre los registros, luego sobre los datos de cada registro.
		`
	fmt.Printf("%s", multi_line)

	xs1 := []string{"Batman", "Jefe", "Vestido de negro"}
	xs2 := []string{"Robin", "Ayudante", "Vestido de colores"}

	fmt.Println(xs1)
	fmt.Println(xs2)

	xss := [][]string{xs1, xs2}
	fmt.Println(xss)

	for i, reg := range xss {
		fmt.Println("Registro:", i)
		for j, val := range reg {
			fmt.Printf("\tÍndice: %v\tvalor: %v\n", j, val)
		}
	}
}

// func Clase3908() {

// 	multi_line := `Ejercicio Práctico #8
// 	Crea un mapa con una llave TIPO string el cual representa el “nombre_apellido” de una persona y un valor de
// 	TIPO []string el cual almacena sus cosas favoritas. Almacena tres registros en tu mapa. Imprime todos sus
// 	valores con su índice de posición en el slice.

// 	eduar_tua, computadoras, montaña, playa
// 	carlos_ramon, leer,comprar, música
// 	juan_bimba, helado, pintar, bailar
// 	`
// 	fmt.Printf("%s", multi_line)

// 	x := map[string][]string{
// 		`eduar_tua`:    []string{`computadoras`, `montañas`, `playa`},
// 		`carlos_ramon`: []string{`leer`, `comprar`, `música`},
// 		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
// 	}

// 	fmt.Println(x)

// 	for llave, valor := range x {
// 		fmt.Println("Registro:", llave)
// 		for i, val := range valor {
// 			fmt.Println("\t", i, val)
// 		}
// 	}
// }

// func Clase3909() {

// 	multi_line := `Ejercicio Práctico #9
// 	Usando el código del ejemplo anterior, agrega un registro a tu mapa, ahora imprime el mapa usando “range”
// 		`
// 	fmt.Printf("%s", multi_line)

// 	x := map[string][]string{
// 		`eduar_tua`:    []string{`computadoras`, `montañas`, `playa`},
// 		`carlos_ramon`: []string{`leer`, `comprar`, `música`},
// 		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
// 	}

// 	// fmt.Println(m)

// 	x[`luis_perez`] = []string{`trabajar`, `playa`, `cerveza`}

// 	for llave, valor := range x {
// 		fmt.Println("Registro:", llave)
// 		for i, val := range valor {
// 			fmt.Println("\t", i, val)
// 		}
// 	}
// }

// func Clase3910() {

// 	multi_line := `Ejercicio Práctico #10
// 	Usando el código del ejemplo anterior, elimina un registro a tu mapa, ahora imprime el mapa usando “range”
// 	`
// 	fmt.Printf("%s", multi_line)

// 	m := map[string][]string{
// 		`eduar_tua`:    []string{`computadoras`, `montaña`, `playa`},
// 		`carlos_ramon`: []string{`leer`, `comprar`, `música`},
// 		`juan_bimba`:   []string{`helado`, `pintar`, `bailar`},
// 	}

// 	//fmt.Println(m)

// 	m[`luis_perez`] = []string{`trabajar`, `playa`, `cerveza`}

// 	delete(m, `juan_bimba`)

// 	for llave, valor := range m {
// 		fmt.Println("Registro:", llave)
// 		for i, val := range valor {
// 			fmt.Println("\t", i, val)
// 		}
// 	}
// }
