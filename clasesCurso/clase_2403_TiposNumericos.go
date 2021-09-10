package clasesCurso

import (
	"fmt"
	"runtime"
)

func Clase2403() {
	/*
			Tipos Numéricos
			●	Enteros
			○	Números sin decimales
		■	También conocidos como enteros
			○	int & uint
			■	“Tamaños de implementación-específica”
			○	Todos los tipos numéricos son diferentes excepto
			■	byte el cual es un alias para uint8
			■	Rune el cual es un alias para int32
			○	Los tipos son únicos
			■	“Esto es un lenguaje de programación estático”
			■	“Las conversiones son requeridas cuando mezclamos diferentes tipos de datos numéricos en una expresión o asignación.
				Por ejemplo, int32 e int no son del mismo tipo aún cuando pueden tener el mismo tamaño en una arquitectura particular.” (fuente)
			○	Regla de oro: solo usa int
			●	floating point
			○	Números con decimales
			■	También conocidos como números reales
			○	Regla de oro: solo usa float64
			●	Overflows
			●	Lectura recomendada - Libro de Caleb Doxsey

	*/
	fmt.Println("Tipos Numéricos")
	x := 42       // asigna int por defecto
	y := 42.34534 // asigna float64 por defecto
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)   // comprueba sistema operativo
	fmt.Println(runtime.GOARCH) // comprueba arquitectura

	/*
		Numeric types
		A numeric type represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are:

		uint8       the set of all unsigned  8-bit integers (0 to 255)
		uint16      the set of all unsigned 16-bit integers (0 to 65535)
		uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
		uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

		int8        the set of all signed  8-bit integers (-128 to 127)
		int16       the set of all signed 16-bit integers (-32768 to 32767)
		int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
		int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

		float32     the set of all IEEE-754 32-bit floating-point numbers
		float64     the set of all IEEE-754 64-bit floating-point numbers

		complex64   the set of all complex numbers with float32 real and imaginary parts
		complex128  the set of all complex numbers with float64 real and imaginary parts

		byte        alias for uint8
		rune        alias for int32

		The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.

		There is also a set of predeclared numeric types with implementation-specific sizes:

		uint     either 32 or 64 bits
		int      same size as uint
		uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

	*/

}
