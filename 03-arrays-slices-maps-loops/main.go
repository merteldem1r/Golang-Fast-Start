package main

import (
	"fmt"
	"time"
)

func main() {
	// *********** Arrays, Slices, Maps, Loops ***********

	// ------------ Arrays:

	var intArr32 [3]int32 // array of 3 integers, default values are 0
	// length of an array is fixed and determined at compile time

	intArr32[0] = 1
	fmt.Println("intArr32:", intArr32)          // [1 0 0]
	fmt.Println("intArr32[1:3]", intArr32[1:3]) // from index 1 to 2 (3 is exclusive), [0 0]

	fmt.Println("&intArr32:", &intArr32[0])   // 0x140000a0004
	fmt.Println("&intArr32[1]", &intArr32[1]) // 0x140000a0008
	fmt.Println("&intArr32[2]", &intArr32[2]) // 0x140000a000c

	// addresses increasing by 4 bytes because int32 takes 4 bytes in memory

	fmt.Println("intArr32 size:", len(intArr32)) // 3

	// initialize array with values
	var intArr [3]int32 = [3]int32{1, 2, 3}
	// or
	// intArr := [3]int32{1, 2, 3}
	fmt.Println("intArr:", intArr) // [1 2 3]

	intArr2 := [...]int32{4, 5, 6}   // compiler can infer the length of the array
	fmt.Println("intArr2:", intArr2) // [4 5 6]

	// ------------ Slices:

	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

	var intSlice []int32 = []int32{1, 2, 3, 4, 5} // slice of integers, length is 5, capacity is 5
	intSlice = append(intSlice, 7)
	fmt.Println("intSlice:", intSlice)               // [1 2 3 4 5 7]
	fmt.Println("intSlice length:", len(intSlice))   // 6
	fmt.Println("intSlice capacity:", cap(intSlice)) // 12 (capacity is doubled when the slice needs to grow)

	var intSlice2 []int32 = []int32{1, 2, 3}
	intSlice = append(intSlice, intSlice2...)       // spreading the elements of intArr2 into intSlice2
	fmt.Println("intSlice after append:", intSlice) // [1 2 3 4 5 7 1 2 3]

	// using make to create a slice with a specific length and capacity
	intSlice3 := make([]int32, 5, 10)                  // length is 5, capacity is 10
	fmt.Println("intSlice3:", intSlice3)               // [0 0 0 0 0]
	fmt.Println("intSlice3 length:", len(intSlice3))   // 5
	fmt.Println("intSlice3 capacity:", cap(intSlice3)) // 10

	// ------------ Maps (key-value pairs):

	var myMap map[string]int = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println("myMap:", myMap)               // map[one:1 three:3 two:2]
	fmt.Println("myMap[four]:", myMap["four"]) // 0 (default value for int)

	// checking if a key exists in the map
	value, ok := myMap["four"]
	if ok {
		fmt.Println("Value for 'four':", value) // Value for 'four': 2
	} else {
		fmt.Println("'four' does not exist in the map")
	}

	// delete value (no return value)
	delete(myMap, "two")
	fmt.Println("myMap after delete:", myMap) // map[one:1 three:3]

	// ------------ Loops:

	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// or

	for i := range 5 {
		fmt.Println("i in range loop:", i)
	}

	var agesMap = map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	for name, age := range agesMap {
		fmt.Println("Name:", name, "Age:", age)

		// Maps implemented as hash tables, so the order of iteration is not guaranteed to be the same every time you run the program. Therefore, the output may vary between runs.
	}

	// for loop
	for i, val := range intSlice {
		fmt.Println("Index:", i, "Value:", val)
	}

	// no while loop in Go, but you can use a for loop to achieve the same result
	i := 0
	for i < len(intSlice) {
		if i == 10 {
			break
		}
		fmt.Println("i in while loop:", i)
		i++
	}

	// ------- TEST (setting capacity of a slice)
	var n int = 1000000
	var withoutCap = []int{}
	var withCap = make([]int, 0, n)

	fmt.Println("Time taken without capacity:", timeLoop(withoutCap, n)) // Time taken without capacity: 9.880167ms
	fmt.Println("Time taken with capacity:", timeLoop(withCap, n))       // Time taken with capacity: 1.837ms
}

func timeLoop(slice []int, n int) time.Duration {
	start := time.Now()
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return time.Since(start)
}
