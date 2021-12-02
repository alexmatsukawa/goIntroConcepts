package main

import (
	"fmt"
	"strconv"
)

//var block; can be used to initialize and organize related variables
var (
	actorName     string = "Jenna Coleman"
	companionName string = "Clara Oswald"
	doctorNumber  int    = 11
	season        int    = 5
)

func main() {
	//i := 0
	var j int = 99 //-> returns 99, int
	//k := 220. //-> returns 220, float64
	//MUST USE DECLARATION IF YOU WANT FLOAT32
	//colon equals syntax cannot be used at package level; must use full declaration
	//NOTE: Local variables must always be used, otherwise it will fail
	//lowercase vars are limited to packages, uppercase vars are global in scope and can be accessed from outside

	//var k float32
	//k = float32(j)//-> type conversion; must use explicit conversion when changing types
	//going from float to int may make you lose information

	var k string
	k = strconv.Itoa(j) //-> needs strconv import to properly convert number to string
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T", k, k)
}
