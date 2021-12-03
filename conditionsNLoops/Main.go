package main

import (
	"fmt"
)

func main() {
	//IF STATEMENTS
	//must use curly braces for if statements
	
	/*
	if true {
		fmt.Println("This test is true")
	}
	*/
	statePopulations := map[string]int{
		"California":	39250017,
		"Texas":		27862596,
		"Florida":		20612439,
		"New York": 	19745289,
		"Pennsylvania":	12802503,
		"Illinois":		12801539,
		"Ohio":			11614373,	
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
		fmt.Println("")
	} else {
		fmt.Println("There is no data for Florida...")
		fmt.Println("")
	}

	//TEST GAME
	number := 50
	guess := 150
	if guess < 1 || guess > 100 {
		fmt.Println("Your guess must be between 1 and 100...")
	} else if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low...")
		} else if guess > number {
			fmt.Println("Too high...")
		} else if guess == number {
			fmt.Println("Correct!")
		}
		fmt.Println(number <= guess, number >= guess, number != guess)
		fmt.Println("")
	}

	//when doing comparison with decimals, don't do direct equivalency comparison
	/*EX: 
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 { --> checks within a hundredth of a percent to see if they are about equal to one another
		fmt.Println("These are the same number")
	} else {
		fmt.Println("These are different numbers")
	}
	*/

	//SWITCH STATEMENTS
	//NOTE: CANNOT HAVE OVERLAPPING SWITCH CHECKS IN SYNTAX BELOW
	switch 6 {
	case 1, 5 , 10:
		fmt.Println("one, five, or ten")
	case 2, 4, 6: 
		fmt.Println("two, four, or six")
	default: 
		fmt.Println("another number")
	}

	//MORE COMPLEX INITIALIZER THAN ABOVE SWITCH CASE
	switch i := 2 + 3; i {
	case 1, 5 , 10:
		fmt.Println("one, five, or ten")
	case 2, 4, 6: 
		fmt.Println("two, four, or six")
	default: 
		fmt.Println("another number")
	}

	//EMPTY SWITCH
	/*
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to 10")
	case i <= 20: 
		fmt.Println("less than or equal to 20")
	default: 
		fmt.Println("greater than twenty")
	}

	Switch case can be left empty if logical comparisons are done on the case statements
	NOTE: IN GO, BREAK KEYWORD IS IMPLIED (i.e. DON'T NEED TO EXPLICITY STATE BREAK FOR EACH CASE)
	Fallthrough keyword will force the program to execute the next case even if the case isn't true

	*/

	//SWITCH CASES AND INTERFACES
	var x interface{} = 1
	switch x.(type) { //-> pulls the actual underlying type of the interface
	case int:
		fmt.Println("x is an int")
		/*
		break -> adding break here will cause program to short circuit and not print what's below this (usage for validation in case)
		fmt.Println("TEST")
		*/
	case float64:
		fmt.Println("x is a float64")
	case string:
		fmt.Println("x is a string")
	default:
		fmt.Println("x is another type")
	}

	//LOOPING
	
}