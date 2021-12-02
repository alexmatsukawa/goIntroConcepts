package main

import (
	"fmt"
)

func main() {
	grades := [...]int{97, 85, 93} //... syntax denotes that array should be only as big enough to carry the data I send in

	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	//fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ray"
	students[2] = "Roald"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("# of Students: %v\n", len(students))
	fmt.Println("")

	//MATRIX EXAMPLE

	//var identityMatrix [3][3]int = [3][3]int { [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1} }
	//another form of declaring the identity matrix
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1,0,0}
	identityMatrix[1] = [3]int{0,1,0}
	identityMatrix[2] = [3]int{0,0,1}
	fmt.Println(identityMatrix)
	fmt.Println("")
	
	//Arrays as Values
	//in go, when you copy the array (line 33), it creates an actual copy of the array instead of just pointing to the original
	a := [...]int{1,2,3}
	b := a
	fmt.Println(a)
	c := &a //-> this points to the array rather than making a copy; means any changes we do on c will reflect on a
	b[1] = 5
	c[2] = 6
	fmt.Println(a)
	fmt.Println(b)
	//NOTE: ARRAY SIZES MUST BE KNOWN AT COMPILE TIME; CANNOT HAVE ANY EXPLICITLY VARIABLE SIZED ARRAYS
	fmt.Println("")
	
	//SLICES
	d := []int{1,2,3} //-> Slices are initialized with empty square brackets
	fmt.Println(d)
	e := d
	e[0] = 2
	fmt.Println(d)
	//everthing we do with arrays, we can do with slices EXCEPT that slices are reference type variables (cannot copy and don't need pointers...)
	fmt.Printf("Length: %v\n", len(d))
	fmt.Printf("Capacity: %v\n", cap(d))
	fmt.Println("")

	//SLICE EXAMPLE
	f := []int{1,2,3,4,5,6,7,8,9,10}
	g := f[:]   //slice of all elements
	h := f[3:]  //slice from 4th element to the end
	i := f[:6]  //slice first 6 elements
	j := f[3:6] //slice elements 4, 5, and 6
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println("")

	//USING MAKE FUNCTION
	x := make([]int, 3, 100) //first param = what type you want, second param = size, third param = capacity
	fmt.Println(x)
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))
	fmt.Println("")

	//USING APPEND (PUSH)
	y := []int{}
	fmt.Println(y)
	fmt.Printf("Length: %v\n", len(y))
	fmt.Printf("Capacity: %v\n", cap(y))
	y = append(y, 9) //first param = source slice, second param = what you want to append
	fmt.Println(y)
	fmt.Printf("Length: %v\n", len(y))
	fmt.Printf("Capacity: %v\n", cap(y))
	y = append(y, 10, 11, 12, 13) 
	fmt.Printf("Length: %v\n", len(y))
	fmt.Printf("Capacity: %v\n", cap(y))
	fmt.Println("")

	//STACK OPERATIONS (POP)
	z := []int{1,2,3,4,5}
	a1 := z[1:] //-> removes first element in the slice
	b1 := z[:len(z) - 1] //-> removes last element in the slice
	fmt.Println(a1)
	fmt.Println(b1)

	c1 := append(z[:2], z[3:]...) //-> removes middle element
	fmt.Println(c1)
	fmt.Println(z) //-> slice z is messed up as a result of the previous operations
	//TO BE COVERED IN LOOPING -> need to create a reference which makes a copy of the initial slice to ensure you have one thing that points to a control copy


}