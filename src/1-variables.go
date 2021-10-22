package main

import "fmt"

// VARIABLES DECLATARION TYPES
func main() {
	// NORMAL VARIABLE DECLATARION
	var normalVar = 1
	fmt.Println(normalVar) // OUTPUTS 1

	// NORMAL CONST DECLARATION
	const normalConst = 1.14
	fmt.Println(normalConst) // OUTPUTS 1.14

	// MULTIPLE VAR DECLARATIONS
	var firstVar, secondVar = 1, 2
	fmt.Println(firstVar, secondVar) // OUTPUTS 1 2

	// DECLARATION SPECIFYING VARIABLE TYPE
	var typedVar int = 2
	fmt.Println(typedVar)

	// SHORT VARIABLE DECLARATION
	shortVar := 5
	fmt.Println(shortVar)
}
