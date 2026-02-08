package main

import (
	"fmt"
	"strings"
)

func main() {
	// *********** Strings, Runes, Bytes ***********
	fmt.Println("Strings, runes, bytes")

	// ----------- Strings:
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed) // 114, uint8

	for i, v := range myString {
		fmt.Println(i, v)
	}
	/*
				r		   é		           S		 U		  m		     é
			[01110010, 11000011, 10101001, 01110011, 01110101, 01101101, 11000011, 10101001]
		    	0 114	1 233					3 115		4 117		5 109		6 233

				it skips the byte at index 2 because it's part of the multi-byte character 'é'
	*/

	/*
		How go reprsents strings in memory:

		UTF-8 encoding is used to represent characters in a string. Each character can be represented by one or more bytes. For example, the character 'é' is represented by two bytes in UTF-8 encoding. When we access a string using an index, we are accessing the byte at that index, not the character. This is why we get 114 (the byte value of 'r') when we access myString[0], and 233 (the byte value of 'é') when we access myString[1].
	*/

	fmt.Printf("The length of 'myString' is %v\n", len(myString)) // 8 (number of bytes, not characters)

	// ----------- Runes:
	var myRune rune = 'é'                  // a rune is an alias for int32, it represents a Unicode code point
	fmt.Printf("%v, %T\n", myRune, myRune) // 233, int32

	var myString2 = []rune("résumé")             // converting string to slice of runes, each rune represents a character
	fmt.Printf("%v, %T\n", myString2, myString2) // [114 233 115 117 109 233], []int32
	var indexedRune rune = myString2[1]
	fmt.Printf("%v, %T\n", indexedRune, indexedRune) // 233, int32

	// ----------- String building:

	// String in golang is immutable, which means that when we concatenate strings, a new string is created in memory. This can be inefficient if we are concatenating a large number of strings in a loop. In such cases, it's better to use a byte buffer or a string builder to build the string more efficiently.

	var strSlice = []string{"m", "e", "r", "t"}
	var castStr string = ""

	// Wrong way to build a string:
	for i := range strSlice {
		castStr += strSlice[i] // inefficient way to build a string, creates a new string in each iteration which can lead to high memory usage and slow performance
	}

	// castStr[0] = 'M' // cannot assign to castStr[0] because strings are immutable, this will cause a compile-time error

	fmt.Println("castStr:", castStr) // "mert"

	// Better way to build a string using a byte buffer:
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // more efficient way to build a string, it uses a buffer to build the string without creating intermediate strings
	}
	var builtStr string = strBuilder.String()
	fmt.Println("builtStr:", builtStr) // "mert"

	// ----------- Bytes:
	var byteSlice = []byte("hello")
	fmt.Printf("%v, %T\n", byteSlice, byteSlice) // [104 101 108 108 111], []uint8
	var indexedByte byte = byteSlice[0]
	fmt.Printf("%v, %T\n", indexedByte, indexedByte) // 104, uint8

	// We can convert a byte slice back to a string:
	var strFromBytes string = string(byteSlice)
	fmt.Printf("%v, %T\n", strFromBytes, strFromBytes) // "hello", string
}
