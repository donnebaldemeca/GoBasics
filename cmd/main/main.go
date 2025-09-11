// File must be UTF-8 encoded
package main

import (
	"errors"
	"fmt"
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
}

// Functions
// keyword functionName(parameterName parameterType) returnType { ... }
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
