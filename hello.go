// name of package
package main

import (
	"errors"
	"fmt"
	"math"
)

// go run     -> Run the program
// go build   -> Compile into executable -> Binary file (.exe)
// go install -> Executable in ../../bin folder -> Compoile outrside dependencies into a package folder

func main() {
	fmt.Println("helllll o0 bitchesjosjsos");
	var x int = 1 // zero value is 0
	var y int = 2
	var sum int = x + y

	// Short syntax
	z := 20   
	fmt.Println(z)

	if sum > 2 { 
		fmt.Println("Greater than 2 biotch")
	} else { 
		fmt.Println("NOt greater than 2 bitch")
	}

	fmt.Println(sum)

	// Array length has to be fixed
	var a [5]int // All empty valiues are filld with its Element-zero-value, in this case 0
	a[2] = 7

	// var aList [3]int{1, 2, 3, 4}
	aList := [3]int{1,2,3} // The length of the array is part of its type
	// To work around this use Slices instead, which don't have a fix # of elements
	aSlice := []int{1,2,3,4} 

	// Add new elemnt
	// .append() creates a new slice, it doesn't modify it, as it created a new one insetead
	aSlice = append(aSlice, 100)

	fmt.Println(aSlice)

	fmt.Println(aList)

	fmt.Println(a)

	// Maps, [KEY_TYPE]VALUE_TYPE
	colors := make(map[string]int)
	colors["red"] = 100
	colors["blue"] = 200
	colors["black"] = 500

	colorRed := colors["red"]

	// Delete from pmap
	delete(colors, "blue")

	fmt.Println(colors)
	fmt.Println(colorRed)

	// LoOps
	// Only one type for...

	for i := 0; i < 5; i++ { 
		fmt.Println(i)
	}

	// While behavior
	j := 10

	for j < 15 { 
		j++
		fmt.Println(j)
	}

	// Iterate over each element
	dummyList := []int{100, 200, 300}
	for index, value := range dummyList { 
		fmt.Println("index:", index, "value:", value) // spaces infered かこいい!
	}

	// Iterate a map
	dummyMap := make(map[string]int)
	dummyMap["a"] = 111
	dummyMap["b"] = 222

	fmt.Println("Mapp looop")
	for key, value := range dummyMap { 
		fmt.Println("key:", key, "value:", value)
		fmt.Println("dummyMap['", key, "']:", dummyMap[key])
	}

	sumResult := addition(100, 200)
	fmt.Println(sumResult)

	// Multiple variable function return
	var input float64 = 100
	result, err := sqrt(input)
	if err != nil { 
		fmt.Println("error:", err)
	} else { 
		// Spaces infered only with Println
		fmt.Println("Square root of", input, "is:", result)
	}

	person := person{name: "Leonardo", age: 23}
	fmt.Println(person)
	fmt.Print("name: ", person.name, " | age: ", person.age)

	// Pointers (memory hehe)
	// This will NOT update xValue's value as it's passing a COPY of xValue
	// Sooo, increase(:int) takes a COPY of it, not its reference in memory (pointer)
	// It only updates the COPIED memory-value, not the one passed
	xValue := 100

    // As increase(:*int) now takes a POINTER instead of a VALUE'
	// this call cannot work no more
	// increase(xValue) 
	fmt.Println(xValue)

	// PASSING A POINTER (DEREFERENCING)
	increase(&xValue)

	// Access memory with & just like in C
	fmt.Print("xValue: ", xValue, " | xValue(Pointer): ", &xValue)
}

func increase(x *int) { 
	*x++ // Access the memory @ pointer
}

func addition(x int, y int) int { 
	return x + y
}

// nil is used just as in Swift or None in python
func sqrt(x float64) (float64, error) { 
	if x < 0 { return 0, errors.New("Number is negative :(") }
	return math.Sqrt(x), nil
}

// Structs
type person struct { 
	name string
	age int
}