// File must be UTF-8 encoded
package main

import (
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
		When variables are initialized without a value default values are set
		Variable types with 0 as default value:
		int, uint, float, complex, byte, rune

		When ini
	*/
}
