package main

import (
	"fmt"
	//"math"
)

const a int16 = 25
const ( //const block
	e = iota //iota is a counter for when we make enumerated constants
	//f = iota
	//g = iota
	f //compiler auto infers pattern; this is the same as commented declarations above
	g
	//NOTE: iota is always scoped to a constant block (i.e. another const block with iota in it will restart at 0)
	
)

const ( //EX: BYTE SIZES CONST BLOCK
	//NOTE: can change starting value of iota by performing operations with it (i.e iota + 5 or iota * 5)
	_ = iota //ignores the initial iota value by assigning to a blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
	//NOTE: iota is always scoped to a constant block (i.e. another const block with iota in it will restart at 0)
	//NOTE: can change starting value of iota by performing operations with it (i.e iota + 5 or iota * 5)
)

func main() {
	//CONST INTRO
	const myConst int = 42 // -> constant we don't want to export starts lowercase
	const a int = 14 // -> this constant overwrites the package level const
	//const MyConst -> constant we do want to export starts uppercase
	//myConst = 22 -> throws error since constants cannot be reassigned
	//const myConst2 float64 = math.Sin(1.57) -> throws error since constants have to be assignable at compile time
	//constants can be instantiated with any primitive types
	//can perform operations on consts and vars as long as they have the same type
	fmt.Printf("%v, %T\n", myConst, myConst)
	fmt.Printf("%v, %T\n", a, a)

	const c = 66
	var d int16 = 24
	fmt.Printf("%v, %T\n", c + d, c + d) // implicit conversion from const allows us to have variable typing for the const value whenever it is used

	//ENUMERATED CONSTS
	fmt.Printf("%v, %T\n", e, e) // result -> 0, int (default instantiation for iota)
	fmt.Printf("%v, %T\n", f, f) // result -> 1, int
	fmt.Printf("%v, %T\n", g, g) // result -> 2, int

	filesize := 4000000000.
	fmt.Printf("%.2f GB\n", filesize/GB)

}