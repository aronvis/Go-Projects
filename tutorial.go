package main

import (
	"fmt"
	"strings"
	"regexp"
	// Rand cannot be accesed through "math" using math.rand
	"math/rand"
	"time"
)

// -------- Variables --------
// Variable Declaration (stores zero value of typeName by default)
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

// Type Casting (explicit only)
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

// Returns tuple explicitly
func swap1(x int, y int) (int, int) {
	return y, x
}

// Retuns values using pointers 
func swap2(x* int, y* int) {
	var temp int = *x
	*x = *y
	*y = temp
}

// Returns tuple implicitly (Should only be used in short functions)
func half(x int, y int) (a, b int) {
	sum := x + y
	a = sum/2
	b = sum/2
	return 
}

// ------ Control Flow -------
// For loop
func forLoop() {
	for i := 0; i<10; i++ {
		fmt.Println(i)
	}
}

// While loop
func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// Infinite Loop
func infLoop() {
	i := 0
	for {
		i += 1
		if i == 10 {
			break
		}
	}
	fmt.Println(i)
}

// If/else Statement
func ifElse() {
	// Remove deterministicness of rand functions
	rand.Seed(time.Now().UnixNano());
	// Optional: Declare 1 var on same line as if statement
	if value := rand.Intn(10);  value % 2 == 0 {
		fmt.Println(value, "Is an even value")
	} else {
		fmt.Println(value, "Is an odd value")
	}

}

// Switch Statement (Terminates when a case succeeds)
func switchStatement() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// Defer Statement
// Statement that won't get executed until the return section of the function is reached
// When multiple defer statements are used, they are ordered using a stack (reserve execution)
func deferStatement() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

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
	forLoop()
	whileLoop()
	infLoop()
	ifElse()
	switchStatement()
	deferStatement()
}