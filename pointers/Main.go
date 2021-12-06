package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	//POINTERS
	a := 42
	b := a //-> this is a copy of a; doesn't actually point to the same thing
	//(if a changes, b will not change)
	fmt.Println(a, b)
	a = 28
	fmt.Println(a, b)

	//We can change the above behavior by using pointers instead
	var c int = 42
	var d *int = &c //-> pointer declaration
	fmt.Println(c, d) //d will print out the memory address of c
	fmt.Println(c, *d) //-> this asterisk is a dereferencing operator
	fmt.Println("")
	//both c and d are linked; if changes are done to one, they will be reflected on the other

	//POINTER ARITHMETIC
	f := [3]int{1,2,3}
	g := &f[0]
	h := &f[1]
	fmt.Printf("%v %p %p\n", f, g, h) //--> g and h will print out the memory location
	fmt.Println("")
	//due to design choices for GO, pointer arithmetic was left out of the language

	//POINTERS AND STRUCTS
	var ms *myStruct
	fmt.Println(ms) //result = nil
	ms = &myStruct{foo:42}
	//can also use the new function to use the above pointer reference
	//ms = new(myStruct) //result = &{0}
	fmt.Println(ms) //result = &{42}; implies that ms is holding the address of an object that has a field of 42 in it

	//To access the underlying struct and it's fields, we have to dereference the pointer.
	(*ms).foo = 55//dereference operator has a lower precedence than the dot operator, so we need parentheses to make sure we're dereferencing ms and not ms.foo
	fmt.Println((*ms).foo)
	//we can omit the parentheses and the dereferencing operator to do the same as above thanks to the GO compiler
	//EX:
	/*
	ms.foo = 55
	fmt.Println(ms.foo)
	*/
	fmt.Println("")
}