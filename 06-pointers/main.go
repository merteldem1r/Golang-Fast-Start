package main

import (
	"fmt"
)

func main() {
	// *********** Pointers ***********
	fmt.Println("Pointers")

	// A pointer is a variable that holds the memory address of another variable. Pointers are used to pass variables by reference, which allows us to modify the original variable from within a function.

	var p *int32                    // declaring a pointer to an int32, default value is nil
	fmt.Printf("p: %v, %T\n", p, p) // <nil>, *int32
	var i int32 = 42
	p = &i                               // assigning the address of i to p
	fmt.Printf("p: %v, *p: %v\n", p, *p) // 0x14000106028, 42

	// using new keyword to create a pointer:
	var p2 *int32 = new(int32)               // creating a new pointer to an int32, default value is 0
	fmt.Printf("p2: %v, *p2: %v\n", p2, *p2) // 0x1400010e02c, 0
	*p2 = 100                                // assigning a value to the pointer
	fmt.Printf("p2: %v, *p2: %v\n", p2, *p2) // 0x1400010e02c, 100

	// invalid memory address or nil pointer dereference is a runtime error that occurs when we try to dereference a pointer that is nil (i.e., it does not point to any valid memory address). This can happen if we declare a pointer but do not initialize it, or if we set a pointer to nil and then try to dereference it.

	// var p3 *int32
	// fmt.Printf("p3: %v, *p3: %v\n", p3, *p3) // <nil>, <nil> // dereferencing a nil pointer will cause a runtime panic

	// using pointers to modify the original variable from within a function:
	var x int32 = 10
	fmt.Printf("Before: x: %v\n", x)
	modify(&x)
	fmt.Printf("After: x: %v\n", x)

	// Pointers on slices:
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice // this creates a copy of the slice header, but both slice and sliceCopy point to the same underlying array
	sliceCopy[0] = 100
	fmt.Printf("slice: %v, sliceCopy: %v\n", slice, sliceCopy) // [100 2 3], [100 2 3] // modifying sliceCopy modifies the underlying array, which is reflected in slice as well

	// so basically we copying the pointer to the underlying array, not the array itself, which is why modifying sliceCopy modifies the original slice as well

	// Pointers on arrays:
	var thing1 = [5]float64{1, 2, 3, 4, 5}

	// Pass by value (copying the entire array):
	fmt.Printf("\nThe memory location of the thingg1 array is: %p\n", &thing1) // 0x14000132000
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v\n", result)                                      // [1 4 9 16 25]
	fmt.Printf("\nThe values of thing1 after calling square function: %v\n", thing1) // [1 2 3 4 5]

	// Pass by reference (passing a pointer to the array):
	var result2 [5]float64 = squarePointer(&thing1)
	fmt.Printf("\nThe result2 is: %v\n", result2)
	fmt.Printf("\nThe values of thing1 after calling squarePointer function: %v\n", thing1) // [1 4 9 16 25] // modifying the array through the pointer modifies the original array as well
}

func modify(x *int32) {
	*x = 20
}

// Memory addresses of the arrays are different because when we pass an array to a function, it is passed by value, which means that a copy of the array is created in memory for the function.
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p\n", &thing2) // 0x14000132030
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

// taking a pointer to an array as a parameter to avoid copying the entire array:
func squarePointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p\n", thing2) // 0x14000132000 (same as thing1)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}
