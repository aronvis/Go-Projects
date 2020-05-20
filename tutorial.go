package main

import (
	"fmt"
)

// -------- Variables --------
// Variable Declaration
var i int

// Variable Initialization
var x int = 5
var a,b int = 2,3
var c,d = 2.4, "Yes"

// Variable Initialization with implicit type 
// e := true - ONLY POSSIBLE INSIDE func()

// Primitive Variable Types
var (
	// ---- Numerical ---- 
	// Signed int (int - int64)
	intValue int64 = -100000000000
	// Unsigned int (uint - uint64)
	uIntValue uint32 = 1 << 32 - 1
	// float32, float64
	floatValue = 2.5

	// ------ Bool -------
	isSunny bool = true

	// ------ Char (byte == uint8) -------
	firstInitial byte = 'a'

	// ------ String ------
	// rune == int32 (unicode code point)
)

// Special Variable Types
const Pi = 3.14
var p *int = &x

// Type Casting
var f = float32(a)

// ------- Functions -------  
// Void
func printFunc() {
	// Contains no default spacing between items
	fmt.Print("This is the value of x:", x, "\n")
	fmt.Println(f)
	e := true 
	fmt.Println(e)
}

// Returns int
func add(x int, y int) int {
	return x + y
}

// Returns tuple implicitly 
func swap(x int, y int) (int, int) {
	return y, x
}

// Returns value explicitly
func half(x int, y int) (a, b int) {
	sum := x + y
	a = sum/2
	b = sum/2
	return 
}


// Classes


func main() {
	printFunc()
	c = float64(add(a, b))
	fmt.Println(c)
	var largeNumb int = 1000
	var smallNumb int = 1 
	largeNumb, smallNumb = half(largeNumb, smallNumb)
	fmt.Println(largeNumb, smallNumb)




}