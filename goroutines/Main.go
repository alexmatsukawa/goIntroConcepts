package main

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

var wg = sync.WaitGroup{} //-> wait group initializer
var counter = 0
var m = sync.RWMutex{} //mutex initializer

func main() {
	//go sayHello() -> doesn't print out because the main function exits faster than this go routine can work 
	//could add a time.Sleep() call to get the Hello to print out, but it is a bad practice
	
	//Better way to get the go routine to work:
	/*
	var msg = " Hello"
	go func() { //-> use anonymous func
		fmt.Println(msg)//potential problem: links the var msg to this func which is on a different execution stack
	} ()
	msg = "goodbye" --> this will execute before the goroutine has a chance to run (race condition); means goodbye will print out from goroutine instead of hello
	*/

	//Better solution to above problem:
	/*
	wg.Add(1) //this tells the wait group to synchronize to below goroutine
	go func(msg string) { 
		fmt.Println(msg)
		wg.Done()
	} (msg) //--> pass in msg by value (i.e. make a copy of the msg var) instead of directly calling it
	time.Sleep(100 * time.Millisecond) -> this is still bad practice since we don't want to link our wait timings to the real world clock
	//Solution? use a wait group
	wg.Wait() -> this allows us to only use just enough time to wait for the goroutine's execution and not rely on the real world clock to wait for us
	*/

	//using multiple goroutines
	/*
	for i := 0; i < 10; i ++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait() 
	go routines are not synced together and are actually racing against each other which is why the console print is always randomized right now
	to fix this, we have to find a way to sync these goroutines together
	*/

	//Solution? Mutex locks
	//runtime.GOMAXPROCS(100) -> forces 100 threads to be open
	//runtime.GOMAXPROCS(1) -> forces only a single thread to be open; this is useful for forcing concurrency w/ no parallelism
	for i := 0; i < 10; i ++ {
		wg.Add(2)
		m.RLock() //-> moving the locks up here prevents the goroutines from racing against each other
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait() 
	//Still a problem -> numbers are not random, but still have multiple repeated numbers
	//goroutines are still racing against each other; to fix this, we have to run a mutex lock outside of the context of the goroutines
	//PROBLEM: with the way we have the progrma set up here, we have destroyed parallelism and concurrency in the application (mutexes force the application to be synchronous and run in a single-threaded way)
	//We have removed the potential ansynchronous benefits of using goroutines by using mutexes...
	fmt.Printf("# of Threads: %v\n", runtime.GOMAXPROCS(-1)) //-> returns the number of OS threads equal to the number of cores available in the machine 
	//General Rule of Thumb for # of threads we want open
	// we want a number of threads equal to the number of cores we have exposed to the application as a minimum
	//develop with GOMAXPROCS value greater than one, but as you get closer to deployment, run with varying values to see which # of threads will perform the best
	
	//BEST PRACTICES
	/*
	* Don't create goroutines in libraries
	*	-> Let the consumer control the concurrency of their program
	*
	* When creating a goroutine, know how it will end
	*	-> Avoid subtle memory leaks
	*
	* Check for race conditions @ compile time
	*	-> add '-race' flag to the go run command to check for race conditions upon compilation
	*/
}

func sayHello() {
	//m.RLock() //block concurrent reading to protect data
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() //unlocks the above mutex lock
	wg.Done()
}

func increment() {
	//m.Lock() //prevents multiple counter increments from going on
	counter++
	m.Unlock()
	wg.Done()
}