package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("Golang Start")

	// *********** Variables ***********

	// -------- Number:

	// var intNum int16
	// var intNum int64
	var intNum int32
	intNum = 5

	var floatNum float32 = 4.24

	fmt.Println("intNum:", intNum)
	fmt.Println("intNum:", floatNum)

	// operations
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// var result float32 = floatNum32 + float32(intNum32) will throw an error

	// type casting
	var result = floatNum32 + float32(intNum32)
	fmt.Println("result:", result) // 12.1

	// -------- Strings:

	var myString string = "Hello World"
	fmt.Println("myString:", myString)

	// concate strings
	myString = "My Name is" + " " + "Mert"
	fmt.Println("myString:", myString)

	// length of a string in golang gives the number of bytes, not characters
	fmt.Println("myString length:", len(myString)) // 15
	fmt.Println("ç length:", len("ç"))             // 2
	// go uses UTF-8 encoding, which means that some characters (like 'ç') can take more than one byte. Therefore, the length of the string in bytes may not always correspond to the number of characters.

	// to truly get the number of characters, you can use the utf8 package
	fmt.Println("çöü", len("çöü"))                    // 6
	fmt.Println("çöü", utf8.RuneCountInString("çöü")) // 3 (correctly counts the number of characters)

	// -------- Characters:

	var myChar rune = 'A'          // rune is an alias for int32 and represents a Unicode code point
	fmt.Println("myChar:", myChar) // 65 (the Unicode code point for 'A')

	// -------- Booleans:

	var isTrue bool = true
	fmt.Println("isTrue:", isTrue) // true

	isFalse := false
	fmt.Println("isFalse:", isFalse) // false

	// -------- Default Values:

	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool

	fmt.Println("defaultInt:", defaultInt)       // 0
	fmt.Println("defaultFloat:", defaultFloat)   // 0
	fmt.Println("defaultString:", defaultString) // ""
	fmt.Println("defaultBool:", defaultBool)     // false

	// dropping the var keyword and type will cause an error

	myNum, myNum2 := 10, 20
	fmt.Println("myNum:", myNum)   // 10
	fmt.Println("myNum2:", myNum2) // 20

	// adding type when it's not obvious can help with readability and prevent errors
	myVar := foo()
	fmt.Println("myVar:", myVar) // 42

	var myVar2 int = foo()         // more readable and prevents errors if foo() changes its return type in the future
	fmt.Println("myVar2:", myVar2) // 42

	// *********** Constants ***********

	const pi float32 = 3.14
	// myConst = 200 // this will throw an error because myConst is a constant and cannot be reassigned
	// myConst2 := 200 // this is not a constant, it's a variable
	fmt.Println("myConst:", pi) // 100

}

func foo() int {
	return 42
}
