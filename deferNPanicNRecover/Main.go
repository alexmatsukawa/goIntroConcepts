package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//DEFER
	//fmt.Println("start")
	//defer fmt.Println("middle") -> defer keyword executes any functions with the key word after all other functions have completed their run times but before the main function returns
	//fmt.Println("end")
	//deferred functions are executed in LIFO (last in first out) order

	//PRACTICAL EXAMPLE
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //-> allows us to associate open and close next to each other while having close still execute at the end
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	//DEFER EXECUTION
	/*
	a := "start"
	defer fmt.Println(a)
	a = "end"

	// defer function takes in the param at the time the defer is called,
	// so this code block prints start instead of end
	*/

	//PANIC
	//In go, you usually throw errors and not exceptions (no explicit exception handling)
	/*
	a, b := 1, 0
	ans := a/b
	fmt.Println(ans) // console will return a panic log

	fmt.Println("start")
	defer fmt.Println("this was deferred") //this will print ot console before panic since defer functions happen before panic happens
	panic("something went wrong...") //instantly throws an error and panic log in console; ends program
	fmt.Println("end") //this line is unreachable due to above panic
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		panic(error.Error()) //ends program with panic upon trying to run a duplicate of this
	}

	//RECOVER
	/*
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil { //recover will return the rror that is causing the program to panic; if program is not panicking, it will return nil
			log.Println("Error:", err)
		}
	}
	panic("something went wrong...") 
	fmt.Println("end")
	*/
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error", err)
			panic(err)
		}
	}()
	panic("something bad happened...")
	fmt.Println("done panicking")
}