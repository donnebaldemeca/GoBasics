// File must be UTF-8 encoded
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	/*
		Variables are statically & strongly typed

		Declaration phrasing:
		keyword theName type
	*/
	var myVariable string // Explicitly declare a variable of type string

	var myNumber = 42         // Declare and initialize a variable, type inferred as int
	myVariable = "Hello, Go!" // Assign a value to the variable

	// Printing using fmt package
	fmt.Println(myVariable)
	fmt.Println("My Number: ", myNumber)

	// Data types
	var intNum int = -8 // integers can be int8, int16, int32, int64 depending on how many bits
	// int by default is int32 on 32-bit systems and int64 on 64-bit systems

	var uintNum uint = 4 // unsigned integers (only positive values), can be uint8, uint16, uint32, uint64

	var float64Num float64 = 12345678.9 // floating-point numbers must be specified as float32 or float64
	var float32Num float32 = 12345678.9 // does not always print as the assigned value, but rather how it is stored in memory due to precision, but float64 is more precise

	fmt.Println("int:", intNum, "\nuint:", uintNum, "\nfloat64:", float64Num, "\nfloat32:", float32Num)

	/*
		Choose the right type based on the need for precision and memory usage
		Cannot perform arithmetic operations between different types without explicit casting
	*/

	var myString string = "Hello, World!" // single line string
	var myMultilineString string = `Hello,
World!` // backticks for multiline strings
	var myConcatString string = "Hello," + " " + "World!" // string concatenation

	fmt.Println(myString)
	fmt.Println(myMultilineString)
	fmt.Println(myConcatString)

	fmt.Println(len("γ"))                    // length of string in bytes, this is lowercase gamma appears as 2 bytes
	fmt.Println(utf8.RuneCountInString("γ")) // length of string in runes(characters), this is 1 rune

	var myRune rune = 'γ'       // rune is an alias for int32, represents a Unicode decimal code or HTML entity (&#947) of a Unicode character
	fmt.Println(myRune)         // prints the Unicode decimal code value1 of the rune
	fmt.Println(string(myRune)) // converts the rune back to a string and prints the character

	var myBoolean bool = true // boolean type
	fmt.Println(myBoolean)

	var myByte byte = 255 // byte is an alias for uint8, represents a single byte of data (0-255)
	fmt.Println(myByte)

	var myComplex complex64 = 1 + 2i // complex numbers, can be complex64 or complex128
	fmt.Println(myComplex)

	var myPointer *string = &myVariable // pointer type, holds the memory address of a variable
	fmt.Println(myPointer)              // prints the memory address of myVariable
	fmt.Println(*myPointer)             // dereference the pointer to get the value of myVariable
	fmt.Println(&myVariable)            // prints the memory address of myVariable

	var myPointerX *string
	fmt.Println(myPointerX) // prints <nil> because the pointer is not initialized
	/*
		When variables are initialized without a value default values are assigned

		Variable types with 0 as default value:
		int, uint, float, complex, byte, rune

		Default value for string is ""
		Default value for bool is false
		Default value for pointer is nil
	*/

	var myInferedInt = 42           // type inferred as int without explicit type declaration
	myInferedString := "Hello, Go!" // shorthand for declaring and initializing a variable, type inferred
	fmt.Println(myInferedInt, myInferedString)

	var1, var2 := "Hello", 42 // multiple variable declaration and initialization
	fmt.Println(var1, var2)

	/*
		Can use short-hand declaration when type is obvious, otherwise use explicit declaration
		Should only use short-hand declaration inside functions, not for package-level variables
	*/

	const myConst string = "constant value" // constant variable, cannot be changed after declaration
	fmt.Println(myConst)

	/*
		Constants can be character, string, boolean, or numeric values
		Constants cannot be declared using the := syntax
		Constants must be initialized with a value
		Constants are often used for configuration values that should not change
	*/
	printMyName("Donne")
	var result, remainder, err = intDivision(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else if remainder == 0 {
		fmt.Printf("10 divided by 3 is %v with no remainder\n", result)
	} else {
		fmt.Printf("10 divided by 3 is %v with a remainder of %d\n", result, remainder)
		// %v is a placeholder 'verb' for any value, %d is a placeholder 'verb' for any decimal value base 10
	}

	/*

		Data Structures

	*/

	// Arrays - fixed size, same type, indexable, contiguous memory
	var intArr [3]int32 = [3]int32{1, 2, 3}
	// var intArr = [...]int32{1, 2, 3} // type inferred, size inferred
	// intArr := [...]int32{1, 2, 3} // shorthand declaration, type inferred, size inferred

	fmt.Println(intArr[0])   // access first element
	fmt.Println(intArr[1:3]) // slice from index 1 to 2 (1 is inclusive, 3 is exclusive)

	// Stored in continguous memory locations 32 bits (4 bytes) each,  4 bytes apart
	fmt.Println(&intArr[0]) // access first element memory location
	fmt.Println(&intArr[1]) // access second element memory location
	fmt.Println(&intArr[2]) // access third element memory location

	// Slices - dynamic size, same type, indexable, contiguous memory, wrapper around arrays
	var intSlice []int32 = []int32{1, 2, 3} // type inferred, size dynamic
	// intSlice := []int32{1, 2, 3} // shorthand declaration, type inferred, size dynamic
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice)) // length and capacity of slice

	intSlice = append(intSlice, 4) // append to slice, increases size dynamically
	// creates a new underlying array if the existing array is not large enough to accommodate the new element
	fmt.Println(intSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(intSlice), cap(intSlice)) // length and capacity of slice after append

	var newIntSlice []int32 = []int32{5, 10}
	newIntSlice = append(newIntSlice, intSlice...) // append intSlice to newIntSlice, ... is the spread operator
	fmt.Println(newIntSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(newIntSlice), cap(newIntSlice))

	var newerIntSlice []int32 = make([]int32, 2, 5) // using the make functrion, creates a slice with length 2 and capacity 5
	// if capacity is not specified, it defaults to the length
	fmt.Println(newerIntSlice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(newerIntSlice), cap(newerIntSlice))

	// Maps - key-value pairs, dynamic size, unordered, reference type
	var myMap map[string]uint8 = make(map[string]uint8) // type inferred, size dynamic
	// myMap := map[string]uint8{"Donne":32, "Alice":28, "Bob": 25} // shorthand declaration, initialized with values
	myMap["Donne"] = 32
	myMap["Alice"] = 28
	myMap["Bob"] = 25
	fmt.Println(myMap)

	// Maps also return a boolean indicating if the key exists
	mapKey := "Jake"
	age, exists := myMap[mapKey]
	if exists {
		fmt.Printf("%s is %d years old\n", mapKey, age)
	} else {
		fmt.Printf("%s not found in map\n", mapKey)
	}

	delete(myMap, "Bob") // delete key-value pair from map
	fmt.Println(myMap)
	fmt.Printf("Length: %d\n", len(myMap)) // length of map

	// Loops
	for name := range myMap {
		fmt.Printf("Name: %s, Age: %d\n", name, myMap[name])
	}
	// In maps the order of iteration is not guaranteed to be the same each time
	// Go does not have a while loop, but can use for loop to achieve the same functionality
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	// Same as while i < 5

	/*
		Performance Testing
	*/
	allocationSize := 100000
	var perfSlice1 = []int{}
	var perfSlice2 = make([]int, 0, allocationSize)

	fmt.Printf("Time taken for slice without pre-allocated capacity: %v\n", timeLoop(perfSlice1, allocationSize))
	fmt.Printf("Time taken for slice with pre-allocated capacity: %v\n", timeLoop(perfSlice2, allocationSize))

	/*
		Strings, Runes, and Bytes
		Strings are immutable, cannot change individual characters
		Runes are used to represent Unicode characters
		Bytes are used to represent raw binary data
	*/

	var thisString = "Résumé"                // string with Unicode characters
	var indexed = thisString[0]              // index the first character, returns a byte value
	fmt.Printf("%v, %T\n", indexed, indexed) // prints the byte value and type of the first character
	fmt.Println(string(indexed))             // converts the byte back to a string and prints the character
	for i, v := range thisString {           // range over the string, returns the index and byte value of each character
		fmt.Printf("Index: %d, Value: %v\n", i, v) // using %c results in rune using %v results in byte value
	}
	// Index 2 is skipped because index 1 and 2 are part of the same Unicode character

	var thisRune = 'a'                                       // rune must be enclosed in single quotes
	fmt.Printf("%v, %T, %c\n", thisRune, thisRune, thisRune) // prints the rune value, type, and character

	var strSlice = []string{"s", "t", "r", "i", "n", "g"} // slice of strings
	var concatStr = ""
	for i := range strSlice { // creates a new string every iteration because strings are immutable
		concatStr += strSlice[i]
	}
	fmt.Println(concatStr)

	// More efficient way to concatenate strings using strings.Builder
	var sb strings.Builder
	for i := range strSlice {
		sb.WriteString(strSlice[i])
	}
	var catStr = sb.String() // convert the builder to a string
	fmt.Println(catStr)

}

/*

	Functions

	Declaration phrasing:
	keyword functionName(parameterName parameterType) returnType { ... }

*/

func printMyName(name string) {
	fmt.Println("My name is", name)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator cannot be zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}

func timeLoop(slice []int, n int) time.Duration {
	start := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(start)
}
