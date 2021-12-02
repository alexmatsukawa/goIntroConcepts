package main

import (
	"fmt"
)

func main() {
	//GENERAL INTEGER INFORMATION
	//n := 1 == 1	//-> returns true
	//m := 1 == 2 //-> returns false
	//var o bool //-> initialized to false (default bool value)

	/*  int has no defined size
	*   int8 goes from -128 to 127
	*   int16 goes from -32,768 to 32,767
	*	int32 goes from -2,147,483,648 to 2,147,483,647
	*	int64 goes from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807  
	*	
	*	uint are unsigned integers
	* 	uint8 goes from 0 to 225
	*	uint16 goes from 0 to 35,535
	*	unit32 goes from 0 to 4,294,967,295
	*/
	//fmt.Printf("%v, %T\n", n, n)
	//fmt.Printf("%v, %T\n", m, m)

	//HEXADECMIAL OPERATIONS
	/*
	a := 10 //hexa -> 1010
	b := 3 //hexa -> 0011

	fmt.Println(a & b) //AND OPERATION, result = 0010
	fmt.Println(a | b) //OR OPERATION, result = 1011
	fmt.Println(a ^ b) //EXCLUSIVE OR OPERATION, result = 1001
	fmt.Println(a &^ b) //AND NOT OPERATION, result = 0100
	*/

	//BIT SHIFTING
	/*
    a := 8
	fmt.Println(a << 3) //bit shift a left 3 places, result = 64
	fmt.Println(a >> 3) //bit shift a right 3 places, result = 1
	*/

	//FLOATING POINT TYPES AND OPERATIONS
	/*
	a := 10.2
	b := 3.7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	var n complex64 = 1 + 2i //also complex128
	fmt.Printf("%v, %T\n", real(n), real(n))//returns 1
	fmt.Printf("%v, %T\n", imag(n), imag(n))//returns 2

	var m complex128 = complex(5, 12) //result = 5 + 12i
	fmt.Printf("%v, %T\n", m, m)
	*/

	//TEXT TYPES
	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2]) //returns byte value not string char values
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	fmt.Printf("%v, %T\n", b, b) //returns byte array of the string s

	r := 'a' //this is a rune; just a type alias for int32
	fmt.Printf("%v, %T\n", r, r)

}