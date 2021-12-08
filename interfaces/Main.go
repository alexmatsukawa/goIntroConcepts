package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	//INTERFACES
	var w Writer = ConsoleWriter{} //-> this can be any other type of writer
	w.Write([]byte("Hello World"))
	fmt.Println("")

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
	fmt.Println("")

	//COMPOSING INTERFACES TOGETHER
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write2([]byte("Hello Go, this is a test!"))
	wc.Close()
	fmt.Println("")

	//TYPE CONVERSION
	bwc := wc.(*BufferedWriterCloser) //type conversion
	//as long as the type conversion succeeds, you can work with the variable as the newly converted type
	//bwc := wc.(io.Reader) -> this returns a panic error since Go does not know how to convert our WriterCloser to an io.Reader
	//Sometimes we want this to work though, so we must do the following...
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed...")
	}

	fmt.Println(bwc)
	fmt.Println("")

	//EMPTY INTERFACE
	var myObj interface{} = NewBufferedWriterCloser() //declaration for an empty interface
	//an empty interface is an interface with no method declarations
	//Useful because we can cast any type to an empty interface since it has not associated methods declared with it
	//Problem: can't really do anything with myObj since it has no methods it exposes
	//To solve this, we need to do type conversion or use the reflect package to figure out what myObj actually is...
	if wc2, ok := myObj.(WriterCloser); ok {
		wc2.Write2([]byte("Hello Go, this is a test!"))
		wc2.Close()
	}
	q, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(q)
	} else {
		fmt.Println("Conversion failed...")
	}
	fmt.Println("")

	//INTERFACE TYPE SWITCHES
	var i interface{} = 0
	switch i.(type) {//type casts the interface to a type to determine what the interface type is
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("what is i...?")
	}
	//Usually paired with the empty interface to determine what type the empty interface is and what types you're expecting to receive
	//NOTE: If any interface methods require a pointer receiver, the interface must be implemented with a pointer.
	//If the interface methods only accept value types, you can use either a pointer or a value to implement the interface.  

	//BEST INTERFACE PRACTICES
	/* 
	* use many small interfaces over a few monolithic ones
	*	EX: io.Writer, io.Reader, interface{}
	* 
	* don't export interfaces for types that will be consumed
	* do export interfaces for types that will be used by the package
	* design functions and methods to receive interfaces whenever possible
	*/
}

type Writer interface { //-> should name interfaces based on what they do (BEST PRACTICES)
	//instead of storing variables, we store method declarations in interfaces
	Write([]byte) (int, error) //method to take in a slice of bytes, and return an int and error value
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int //type alias
//Don't necessarily need to have a struct to create an interface; can use any type
//can't add method to int; but on other types that we can control, we can add methods using this interface declaration
func (ic *IntCounter) Increment() int {
	*ic++ //incrementing the type alias
	return int(*ic)
}

type Writer2 interface {
	Write2([]byte) (int, error)
}

type Closer interface {
	Close() error
}

//THIS IS HOW YOU CAN GROUP MULTIPLE INTERFACES TOGETHER
//This will work as long as you remember to implement the methods that are linked to each interface that is attached here
type WriterCloser interface {
	Writer2
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write2(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {//constructor function to initalize the pointer to BufferedWriterCloser
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

