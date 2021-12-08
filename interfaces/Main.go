package main

import (
	"bytes"
	"fmt"
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
	fmt.Println(bwc)
	fmt.Println("")
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