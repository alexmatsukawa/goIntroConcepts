package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number int
	actorName string
	companions []string
}

type Animal struct {
	Name string `required max:"100"` //-> this is a tag
	Origin string
}

type Bird struct{
	Animal //-> This deirectly embeds the Anima struct into the bird struct
	SpeedKPH float32
	CanFly bool
}

func main() {
	//MAPS
	statePopulations := map[string]int{
		"California":	39250017,
		"Texas":		27862596,
		"Florida":		20612439,
		"New York": 	19745289,
		"Pennsylvania":	12802503,
		"Illinois":		12801539,
		"Ohio":			11614373,	
	}
	//NOTE: SLICES, MAPS, and some other functions cannot be used for equivalency checking
	//m := map[[]int]]string{} --> throws error b/c slice cannot be a valid key type
	//m := map[[3]int]]string{} --> works b/c arrays can be used as key types
	fmt.Println(statePopulations)

	//USING MAKE TO CONSTRUCT MAPS
	/*
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":	39250017,
		"Texas":		27862596,
		"Florida":		20612439,
		"New York": 	19745289,
		"Pennsylvania":	12802503,
		"Illinois":		12801539,
		"Ohio":			11614373,	
	}
	fmt.Println(statePopulations)

	THIS IS THE SAME RESULT AS ABOVE
	*/

	//MAP MANIPULATION
	statePopulations["Georgia"] = 10310371 //directly adds a new key-value pair to map
	fmt.Println(statePopulations["Texas"]) //-> returns the value associated with the key
	fmt.Println(statePopulations["Georgia"])
	fmt.Println(statePopulations) // NOTE: Map return value is not always guaranteed to be the same
	delete(statePopulations, "Georgia") //delete function deletes a key value pair based on key value given
	fmt.Println(statePopulations)
	
	pop, ok := statePopulations["Oho"]
	fmt.Println(pop, ok) //--> ok will return false if key has no associated value located within the map
	fmt.Println(len(statePopulations)) //returns size of map
	//map data is passed in by reference
	sp := statePopulations
	delete(sp, "Florida")
	fmt.Println(sp)
	fmt.Println(statePopulations)
	fmt.Println("")

	//STRUCTS
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor) //returns all struct data
	fmt.Println(aDoctor.actorName) //returns specific element of the struct
	fmt.Println(aDoctor.companions)
	fmt.Println(aDoctor.companions[2])
	fmt.Println("")

	//IN-LINE STRUCT DECLARATION
	bDoctor := struct{name string}{name: "David Tennant"} //-> anonymous struct; CAN'T BE USED ANYWHERE ELSE
	fmt.Println(bDoctor)

	//NOTE: STRUCTS ARE VALUE TYPES AND NOT REFERENCE TYPES; CHANGES WILL NOT BE MADE TO UNDERLYING COPY
	anotherDoctor := aDoctor
	anotherDoctor.actorName = "Matt Smith"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)

	fmt.Println("")

	//COMPOSITION AND EMBEDDING
	//NOTE: GO DOES NOT SUPPORT INHERITANCE; INSTEAD, GO SUPPORTS COMPOSITION THROUGH EMBEDDING

	b := Bird{}
	b.Name = "Ostrich"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	//Can also declare bird as the following:
	/*
	b := Bird{
		Animal: Animal{Name: "Ostrich", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly: false,
	}
	*/
	fmt.Println(b)
	fmt.Println("")
	//TO USE DATA INTERCHANGEABLY, WE NEED TO USE INTERFACES...

	//USING REFLECT TO GET TAG NAME
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}