package main

import (
	"fmt"
	"math/cmplx"

	"rsc.io/quote"
)

var i, j int = 1, 2

func main() {

	//Hello!
	fmt.Println(quote.Hello())
	fmt.Println(add(3, 4))

	//#1 ---- > VARIABLES
	// variables con inicializadores
	var l, python, java = true, false, "No!"
	fmt.Println(i, j, l, python, java) // ----> 1 2 true false No!

	// :=
	var a, b int = 1, 2
	c := 3
	d, python, java := false, true, "Si!"
	fmt.Println(a, b, c, d, python, java) // ----> 1 2 3 false true Si!

	// Tipos Basicos
	/*bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte 				// alias for uint8

	rune // alias for int32
		// represents a Unicode code point

	float32 float64

	complex64 complex128*/

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}

func add(x, y int) int {
	return x + y
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
