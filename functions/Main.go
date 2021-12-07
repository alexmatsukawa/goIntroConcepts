package main

import (
	"fmt"
)

func main() {
	for i:= 0; i < 5; i++ {
		sayMessage("Hello World!", i)
	}
	fmt.Println("")

	//PASSING IN POINTERS
	greeting := "Hello"
	name := "Stacy"
	sayGreeting(&greeting, &name)
	fmt.Println("")

	//VARIATIC PARAMS
	sum(1, 2, 3, 4, 5)
	fmt.Println("")

	//RETURN VALUES
	s := sum2(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)
	fmt.Println("")

	//RETURNING LOCAL VARIABLES AS POINTERS
	s2 := sum3(1, 2, 3, 4, 5)
	fmt.Println("The sum is", *s2) //-> must dereference the pointer from the function
	fmt.Println("")

	//NAMED RETURN VALUES
	s3 := sum4(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s3)
	fmt.Println("")

	//MULTIPLE RETURN VALUES
	d, err := divide(5.0, 3.0) // needs multiple values when returning multiple values
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	fmt.Println("")

	//FUNCTIONS AS TYPES
	
}

//HELPER FUNCTION EXAMPLE
func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

//FUNCTION EXAMPLE W/ PASSING IN POINTERS
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"//with pointers, we can manipulate the data that is passed into the function
	//EX: name will now be Ted instead of Stacy due to pointer manipulation
	//Why do this? passing in a pointer will allow the function to change and manipulate the params,
	//which is necessary sometimes. Pointers are also more efficient to send in as data since they 
	//will be more condensed forms of data compared to a full object.
	//NOTE: CANNOT PASS IN SLICES AND MAPS AS POINTERS SINCE THEY ALREADY USE INTERNAL POINTERS TO HANDLE THEIR DATA.
	fmt.Println(*name)
}

//FUNCTION EXAMPLE W/ VARIATIC PARAMS
func sum(values ...int) { //this notation tells the function to take in the params and wrap them into a slice
	fmt.Println(values) 
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is", result)
	//NOTE: When using variatic params, you cannot use more than one and it must be the last params in the param list.
}

//FUNCTION EXAMPLE W/ RETURN VALUES
func sum2(values ...int) int { //int in the function signature indicates what the return value type is
	fmt.Println(values) 
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

//FUNCTION EXAMPLE W/ RETURNING LOCAL VARIABLES AS POINTERS
func sum3(values ...int) *int { //return value request an int pointer
	fmt.Println(values) 
	result := 0
	for _, v := range values {
		result += v
	}
	return &result //we're returning an address location of result
	//GOLang automatically promotes the local var 'result' once it sees that it needs to return as a pointer
	//Means that the memory location isn't destroyed upon function completion unlike other languages
}

//FUNCTION EXAMPLE W/ NAMED RETURN VALUES
func sum4(values ...int) (result int) { //second parenthesis indicated an implicit return variable
	fmt.Println(values) 
	for _, v := range values {
		result += v
	}
	return //because of the function signature change, we don't need to initialize or explicitly state result in the function
}

//FUNCTION EXAMPLE W/ MULTIPLE RETURN VALUES
func divide(a, b float64) (float64, error) {//Multiple return values in function signature
	if b == 0.0 {
		//panic("Cannot provide zero as a second value") -> don't want to actually panic because we want program to contnue to run instead of closing
		return 0.0, fmt.Errorf("Cannot divide by zero") //multiple return values
	}
	return a / b, nil //multiple return values
}
