// File must be UTF-8 encoded
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// Go Routine variables
var waitGroup = sync.WaitGroup{} // create a waitgroup to wait for all go routines to finish, essentially a counter
// Add waitgroup.Add(1) before starting a go routine
// Add waitgroup.Wait() in main function to wait for all go routines to finish
// Add waitgroup.Done() at the end of the go routine (within function) to decrement the counter

// Simulated database data
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

// Mutex / Locks
var mutex = sync.Mutex{} // create a mutex to lock access to shared resources in go routines
var dbResults = []string{}

// Go Routine variables end

/*

	Structs and Interfaces
	Declaration phrasing:
	type structName struct {
		fieldName fieldType (values default to zero values if not initialized)
		...
	}

*/

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo engineOwner // fields can be other structs, creating nested structs
}

// func (receiverName receiverType) methodName(parameterName parameterType) returnType { ... }
func (g gasEngine) milesLeft() uint8 { // method with receiver of type gasEngine, directly associated with the struct, and can access its fields
	return g.mpg * g.gallons
}

type engineOwner struct {
	name string
	ownerID
}

type ownerID struct {
	id uint8
}

type electricEngine struct {
	mpkwh     uint8
	kwh       uint8
	ownerInfo engineOwner
}

func (e electricEngine) milesLeft() uint8 { // method with receiver of type electricEngine
	return e.mpkwh * e.kwh
}

type engine interface { // interface type, defines a set of methods that a type must implement to satisfy the interface
	milesLeft() uint8 // any type that has a milesLeft method with this signature satisfies the engine interface
}

func canDrive(e engine, miles uint8) { // function that takes an engine interface as a parameter
	if miles <= e.milesLeft() {
		fmt.Println("You can drive!")
	} else {
		fmt.Println("You need to refuel/recharge!")
	}
}

func main() {
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Variables and Data Types")
	fmt.Println(strings.Repeat("-", 50))
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

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Data Structures")
	fmt.Println(strings.Repeat("-", 50))

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

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Performance Test")
	fmt.Println(strings.Repeat("-", 50))

	/*
		Performance Test
	*/
	allocationSize := 100000
	var perfSlice1 = []int{}
	var perfSlice2 = make([]int, 0, allocationSize)

	fmt.Printf("Time taken for slice without pre-allocated capacity: %v\n", timeLoop(perfSlice1, allocationSize))
	fmt.Printf("Time taken for slice with pre-allocated capacity: %v\n", timeLoop(perfSlice2, allocationSize))

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Strings, Runes, and Bytes")
	fmt.Println(strings.Repeat("-", 50))

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

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Structs, Interfaces, and Methods")
	fmt.Println(strings.Repeat("-", 50))

	// Structs, Interfaces, and Methods
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15} // initialize struct
	// myEngine := gasEngine{25, 15} // shorthand declaration, order matters for field values when initializing without field names
	// if no value is provided for a field, it defaults to the zero value of the field type
	// myEngine.mpg = 25 // can assign values to struct fields individually
	// myEngine.gallons = 15
	fmt.Printf("My engine gets %d miles per gallon and has a %d gallon tank\n", myEngine.mpg, myEngine.gallons)

	myEngine.ownerInfo = engineOwner{name: "Donne", ownerID: ownerID{id: 1}} // initialize nested struct field
	fmt.Printf("Engine owner is %s, ID %d\n", myEngine.ownerInfo.name, myEngine.ownerInfo.id)

	var myInfo engineOwner = engineOwner{name: "Alice", ownerID: ownerID{id: 2}} // initialize nested struct with nested struct field
	fmt.Printf("Owner is %s, ID %d\n", myInfo.name, myInfo.id)                   // can also access nested struct fields directly myInfo.id instead of myInfo.ownerID.id

	// Anonymous struct
	var hydroEngine = struct { // no name for the struct type, cannot be reused
		waterCapacity  uint8
		estimatedRange uint16
		ownerInfo      engineOwner
	}{waterCapacity: 100, estimatedRange: 300, ownerInfo: engineOwner{name: "Bob", ownerID: ownerID{id: 3}}}
	fmt.Printf("Hydro engine has %d gallons of water and an estimated range of %d miles. Owner is %s, ID %d\n", hydroEngine.waterCapacity, hydroEngine.estimatedRange, hydroEngine.ownerInfo.name, hydroEngine.ownerInfo.id)

	myEngine.gallons = 3 // milesLeft() returns uint8, if gallon value was 15 it would overflow and return incorrect value

	fmt.Printf("My gas engine can go %d miles before refueling\n", myEngine.milesLeft()) // call method on struct

	var myElectricEngine electricEngine = electricEngine{mpkwh: 3, kwh: 10, ownerInfo: engineOwner{name: "Eve", ownerID: ownerID{id: 4}}}
	canDrive(myElectricEngine, 50) // pass struct that implements the engine interface

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("Pointers and Memory Management")
	fmt.Println(strings.Repeat("-", 50))

	/*

		Pointers and Memory Management

	*/

	var pointer *int32 = new(int32) // new allocates memory for an int32 and returns a pointer to it
	// pointer will default to nil if not initialized
	// pointer points to empty memory location with size of int32 (4 bytes)
	var integer32 int32 = 3
	fmt.Printf("The value pointer points to is: %v\n", *pointer)       // dereference pointer to get value, defaults to 0
	fmt.Printf("The memory address of pointer is: %v\n", pointer)      // prints memory address
	fmt.Printf("The value of integer32 is: %v\n", integer32)           // defaults to 0
	pointer = &integer32                                               // assign the address of integer32 to pointer
	fmt.Printf("The value pointer points to is: %v\n", *pointer)       // dereference pointer to get value
	fmt.Printf("The memory address of pointer is: %v\n", pointer)      // prints memory address
	fmt.Printf("The memory address of integer32 is: %v\n", &integer32) // prints memory address of integer32
	*pointer = 10                                                      // change the value at the memory address pointer points to
	fmt.Printf("The value pointer points to is: %v\n", *pointer)       // dereference pointer to get value
	fmt.Printf("The value of integer32 is: %v\n", integer32)           // integer32 value has changed to 10

	// Pointers and slices
	// Slices use pointers internally to reference the underlying array
	var exampleSlice = []int{1, 2, 3}
	var sliceCopy = exampleSlice // creates a copy of the slice header, but both slices point to the same underlying array
	sliceCopy[0] = 10            // changing the value of sliceCopy also changes exampleSlice
	fmt.Println("exampleSlice:", exampleSlice)

	// Pointers and functions
	var floatArray = [5]float64{1, 2, 3, 4, 5}
	fmt.Println("Original array:", floatArray)
	squaredArray := squareArray(&floatArray) // passing array by reference using pointer, instead of by value, thus saving memory
	fmt.Println("Squared array returned from function:", squaredArray)
	fmt.Println("Original array after function call:", floatArray) // original array is unchanged

	/*

		Go routines

		Launch multiple threads of execution within a single program

	*/

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitGroup.Add(1) // increment the waitgroup counter before starting a go routine

		// dbCall(i) // sequential calls, takes longer
		go dbCall(i) // concurrent calls, takes less time, use 'go' keyword infront of function
		// go routines run in the background, main function may exit before they complete, so a waitgroup or sleep may be needed to wait for them to finish
		// dbCall function calls waitGroup.Done() to decrement the counter when it completes
	}
	waitGroup.Wait() // wait for all go routines to finish
	fmt.Printf("Sequential DB calls took: %v\n", time.Since(t0))

	// Mutex / Locks
	t1 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitGroup.Add(1)
		go dbCallMutexLock(i)
		// dbCallMutexLock function uses mutex to lock access to shared resource (dbResults slice) when writing to it
	}
	waitGroup.Wait()
	fmt.Printf("Sequential DB calls took: %v\n", time.Since(t1))

	/*

		Channels

		Hold data sent from one go routine to another
		Thread-safe communication between go routines
		Listens for data to be sent to it
		Buffered or unbuffered
		Uses keywords: make, chan, <- (channel operator)

	*/

	var channel = make(chan int) // can only hold a single int value
	// must be used with go routines to send and receive data concurrently

	/*

		this will not work as there is no go routine to receive the value being sent to the channel, thus it will block forever

		channel <- 42                //send value to channel, blocks until a go routine is ready to receive the value
		// channel is uses underlying array of size 1, if another value is sent to the channel before the first is received, it will block

		var chanVar = <-channel // receive value from channel, blocks until a value is sent to the channel
		fmt.Println("Value received from channel:", chanVar)
	*/

	go channelProcess(channel) // start a go routine to send a value to the channel
	var chanVar = <-channel    // receive value from channel, blocks until a value is sent to the channel
	fmt.Println("Value received from channel:", chanVar)

	go channelProcessLoop(channel) // start a go routine to send multiple values to the channel
	for v := range channel {       // receive values from channel until it is closed, blocks until a value is sent to the channel
		fmt.Println("Value received from channel:", v)
	} // prints as values are received from the channel, fast

	// Buffer channels
	var bufferChannel = make(chan int, 5)
	go channelProcessLoop(bufferChannel) // channelProcessLoop function process ends before the receiving loop ends, because the channel has a buffer of 5 and can hold all values sent to it before blocking
	for v := range bufferChannel {
		fmt.Println("Value received from buffered channel:", v)
		time.Sleep(time.Second * 1) // simulate slow processing of received values
	}
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

// Pointers and functions
func squareArray(fl64Value *[5]float64) [5]float64 {
	for i := range fl64Value {
		fl64Value[i] = fl64Value[i] * fl64Value[i]
	}
	return *fl64Value
}

// Go routine function example
func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("DB call %d took %f seconds\n", i, delay/1000)
	waitGroup.Done() // decrement the waitgroup counter when the go routine completes
}

func dbCallMutexLock(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("DB call %d took %f seconds\n", i, delay/1000)

	mutex.Lock() // lock access to shared resource
	// Necessary to prevent threads from writing to the shared resource at the same time

	dbResults = append(dbResults, dbData[i]) // simulate storing result in a shared resource
	mutex.Unlock()                           // lock access to shared resource
	// Can also use Read/Write mutex for more granular control over read and write access
	// Can use Rlock() and RUnlock() for read access

	waitGroup.Done() // decrement the waitgroup counter when the go routine completes
}

// Go routine for channels example

func channelProcess(ch chan int) {
	ch <- 42 // send value to channel
}

func channelProcessLoop(ch chan int) {
	defer close(ch) // closes the channel when the function exits
	// keyword defer delays the execution of a function until the surrounding function returns, last statement to be executed
	for i := 0; i < 5; i++ {
		ch <- i // send value to channel
	}
	fmt.Println("Channel sender done sending values")
	// close(ch) // can also close the channel here, but defer is more reliable
	// closing channel necessary to prevent deadlock when ranging over the channel in the receiving go routine
}
