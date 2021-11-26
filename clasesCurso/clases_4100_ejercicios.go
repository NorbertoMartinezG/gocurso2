package clasesCurso

import (
	"fmt"
)

type persona41 struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func Clase4111() {

	multi_line := `Ejercicio práctico #1
	Crea tu propio tipo “persona” el cual tendrá un tipo subyacente tipo “struct” de manera que pueda 
	almacenar la siguiente data:
	●	Nombre
	●	Apellido
	●	Sabores de helado favoritos
	
	Crea dos VALORES de TIPO persona. Imprime los valores, usa range sobre los elementos en el slice 
	el cual almacena los valores de helados favoritos.
	
		`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	p1 := persona41{
		nombre:   "Eduar",
		apellido: "Tua",
		saboresFav: []string{
			"chocolate",
			"mantecado",
			"torta suiza",
		},
	}
	p2 := persona41{
		nombre:   "Condor",
		apellido: "Pérez",
		saboresFav: []string{
			"fresa",
			"vainilla",
			"limón",
		},
	}
	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)
	for i, v := range p1.saboresFav {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.nombre)
	fmt.Println(p2.apellido)
	for i, v := range p2.saboresFav {
		fmt.Println("\t", i, v)
	}

}

func Clase4112() {

	multi_line := `Ejercicio práctico #2
	Usa el código del ejemplo anterior y almacena los valores de tipo persona en un mapa con la 
	llave apellido. Accede a cada valor en el mapa. Imprime los valores. Imprime también los valores 
	haciendo range sobre el slice.
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	p1 := persona41{
		nombre:   "Eduar",
		apellido: "Tua",
		saboresFav: []string{
			"chocolate",
			"mantecado",
			"torta suiza",
		},
	}
	p2 := persona41{
		nombre:   "Condor",
		apellido: "Pérez",
		saboresFav: []string{
			"fresa",
			"vainilla",
			"limón",
		},
	}

	m := map[string]persona41{
		p1.apellido: p1,
		p2.apellido: p2,
	}

	for _, v := range m {
		fmt.Println(v.nombre)
		fmt.Println(v.apellido)
		for i, v := range v.saboresFav {
			fmt.Println(" ", i, v)
		}
		fmt.Println("-------")
	}

}

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func Clase4113() {

	multi_line := `Ejercicio práctico #3
	●	Crea un nuevo tipo: vehículo. 
		○	El tipo subyacente es un struct. 
		○	Los campos: 
			■	puertas
			■	color 
	●	Crea dos nuevos tipos: camión & sedán. 
		○	El tipo subyacente de cada uno de esos tipo es un struct. 
		○	Embebe el tipo “vehículo” en ambos camión y sedán. 
		○	Dale al camión el campo “cuatroRuedas” el cual será un booleano. 
		○	Dale al sedán el campo “lujoso” el cual será un booleano.
	●	Con los structs vehículo, camión y sedán: 
		○	Usando un composite literal, crea un valor de tipo camión y asígnale valor a los campos. 
		○	Usando un composite literal, crea un valor de tipo sedan y asígnale valor a los campos. 
	●	Imprime cada uno de los valores. 
	●	Imprime un solo valor de cada uno de eso valores.
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	c := camion{
		vehiculo: vehiculo{
			puertas: 2,
			color:   "blanco",
		},
		cuatroRuedas: true,
	}

	s := sedan{
		vehiculo: vehiculo{
			puertas: 4,
			color:   "negro",
		},
		lujoso: false,
	}

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(c.puertas)
	fmt.Println(s.puertas)

}

func Clase4114() {

	multi_line := `Crea y usa un struct anónimo.
	`
	fmt.Printf("%s", multi_line)
	//----------------------------------------------------------------------

	s := struct {
		nombre     string
		amigos     map[string]int
		bebidasFav []string
	}{
		nombre: "Eduar",
		amigos: map[string]int{
			"Carito":  222,
			"Adriana": 444,
			"Condor":  666,
		},
		bebidasFav: []string{
			"agua",
			"naranjada",
			"cerveza",
		},
	}

	fmt.Println(s.nombre)
	fmt.Println("\tAmigos:")
	for k, v := range s.amigos {
		fmt.Println("\t\t", k, v)
	}
	fmt.Println("\tBebidas favoritas:")
	for i, v := range s.bebidasFav {
		fmt.Println("\t\t", i, v)
	}

}
