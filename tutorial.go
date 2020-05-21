package main

import (
	"fmt"
	"strings"
	"regexp"
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

	// ------ Char (runs == ASCII value of char) See below -------
	isCap = isCapital('A')

	// ------ String ------
	fullName = "Aron Vischjager"
	
)

// Special Variable Types
const Pi = 3.14
var p *int = &x

// Type Casting
var f = float32(a)

// -------- String member functions -------
func stringFunc() {
	// Contains
	fmt.Println(strings.Contains("test", "es"))
	// Count
	fmt.Println(strings.Count("test", "t"))
	// Index - Returns first instance
	fmt.Println(strings.Index("tweet", "e"))
	// Join 
	fmt.Println(strings.Join([]string{"a", "b"}, ","))
	// Equal and compare (ignore case)
	fmt.Println(strings.EqualFold("test", "TEST"))
	// Length 
	fmt.Println(len("test"))
	// Replace - Replaces first n instances 
	fmt.Println(strings.Replace("foo", "o", "0", 1))
	// Replace all
	fmt.Println(strings.ReplaceAll("foo", "o", "0"))
	// Split
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	// ToLower
	fmt.Println(strings.ToLower("TEST"))
	// ToUpper
	fmt.Println(strings.ToUpper("test"))
}

// -------- Regexp member functions -------
func regexpFunc() {
	// Returns true if the string matches the pattern
	var match, error = regexp.MatchString(`foo.*`, "seafood")
	fmt.Println("The string matches the pattern:", match, "\nThe following errors occured:", error)
	reg, _ := regexp.Compile("^(The|the).*Spain")
	// Returns true if the string matches the pattern
	fmt.Println(reg.MatchString("The rain in Spain, the rain in the USA, and the rain in spain."))
	// Returns the starting string index of each match
	fmt.Println(reg.FindStringIndex("The rain in Spain, the rain in the USA, and the rain in spain."))
}

// ------- Functions -------  
// Void
func printFunc() {
	// Contains no default spacing between items
	fmt.Print("This is the value of x:", x, "\n")
	fmt.Println(f)
	e := true 
	fmt.Println(e)
}


// Bool
// Used to deal with ASCII/Unicode values (ASCII could also be represented using bytes)
// Rune range: 0 - 2^32 -1 
// Byte range: 0 - 255
func isCapital(r rune) bool {
    switch {
    case 65 <= r && r <= 90:
        return true
    default:
        return false
    }
}

// Returns int
func add(x int, y int) int {
	return x + y
}

// Returns tuple implicitly 
func swap1(x int, y int) (int, int) {
	return y, x
}

// Retuns values using pointers 
func swap2(x* int, y* int) {
	var temp int = *x
	*x = *y
	*y = temp
}

// Returns value explicitly
func half(x int, y int) (a, b int) {
	sum := x + y
	a = sum/2
	b = sum/2
	return 
}

// ------ Classes and Structs ------- 
// ------ Data Structures -------
// ------ Common Algorithms -------
// Sorting a list


func main() {
	printFunc()
	c = float64(add(a, b))
	fmt.Println(c)
	var largeNumb int = 1000
	var smallNumb int = 1 
	largeNumb, smallNumb = half(largeNumb, smallNumb)
	fmt.Println(largeNumb, smallNumb)
	fmt.Println(isCapital('A'))
	fmt.Println(isCapital('a'))
	stringFunc() 
	regexpFunc()
	xValue := 64
	yValue := 8
	swap2(&xValue, &yValue)
	fmt.Println(xValue, yValue)
}