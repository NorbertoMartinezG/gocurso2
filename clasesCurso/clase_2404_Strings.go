package clasesCurso

import "fmt"

func Clase2404() {
	println("heil")

	/*
		Tipo String
		●	“Un valor string es una (posiblemente vacía) secuencia de bytes. Los strings son inmutables: una vez creados, es imposible
			cambiar su contenido.”
		●	En Go, una cadena es, en efecto, un segmento de bytes de solo lectura.
	*/

	s1 := "Hola mundo"
	s2 := `Esta es una línea 
	partida.`

	fmt.Println(s1)
	fmt.Printf("El tipo de s1 es: %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("El tipo de s2 es: %T\n", s2)

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("Imprimir un Slice de bytes")
	bs := []byte(s1) // slice de byte. (conversion de string a un slice de bytes)
	fmt.Println(bs)
	fmt.Println("Imprimir el tipo de la variable bs-> un Slice de bytes")
	fmt.Printf("%T\n", bs) // tipo de dato bs
	fmt.Println("-------------------------------------------------------------------------")

	fmt.Println("")
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("Imprimir el tipo de dato mediante ciclo for el string s1")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}
	fmt.Println("-------------------------------------------------------------------------")

}

func Clase2404_2() {

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("imprimir Unicode de un String")
	s := "Hello, 世界"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) // %#U imprime unicode
	}
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("-------------------------------------------------------------------------")

	s1 := "Hola mundo"
	fmt.Println("ciclo que itera el String -- Hola mundo -- dado su indice y su valor en notacion ")
	for i, v := range s1 {
		fmt.Printf("El índice %d representa el valor %#x\n", i, v) // %#x\n = notacion
	}
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Printf("Con el verbo %q indico que se imprima el string %s", "%s", s1)

	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Println("OTROS EJERCICIOS.")

	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("---%x\n", "世")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	s = "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}
}
