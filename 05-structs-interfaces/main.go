package main

import "fmt"

// STRUCTS: A struct is a composite data type that groups together variables under a single name. Each variable in a struct is called a field. Structs are used to represent complex data structures and can have methods associated with them.

type owner struct {
	name string
}

type gasEngine struct {
	mpg       uint8 // miles per gallon
	gallons   uint8 // how many gallons of fuel the engine has
	ownerInfo owner // owner information of the engine, this is a struct field of type 'owner'
}

type electricEngine struct {
	mpk uint8 // miles per kilowatt-hour
	kwh uint8 // how many kilowatt-hours of charge the engine has
}

// methods of the structs:
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons // calculating how many miles the engine can go with the remaining fuel
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpk * e.kwh // calculating how many miles the engine can go with the remaining charge
}

// INTERFACES: An interface is a type that defines a set of method signatures. A type implements an interface by implementing all the methods defined in the interface. Interfaces are used to achieve polymorphism and to define behavior that can be shared across different types.

type engine interface {
	milesLeft() uint8 // method that calculates how many miles the engine can go with the remaining fuel or charge
}

// using interface to make it more general and take any engine (gas or electric) and check if it can make it to a certain distance:
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You cannot make it!")
	}
}

func main() {
	// *********** Structs, Interfaces ***********

	// ----------- Slices introduction:

	var myEngine gasEngine
	fmt.Printf("myEngine: %v, %T\n", myEngine, myEngine)                                                                               // {0 0}, main.gasEngine
	fmt.Printf("myEngine.mpg: %v, myEngine.gallons: %v, myEngine.ownerInfo: %v\n", myEngine.mpg, myEngine.gallons, myEngine.ownerInfo) // 0, 0, "" // default values of the fields are 0

	// initializing with struct literal:
	var myEngine2 gasEngine = gasEngine{
		mpg:       30,
		gallons:   10,
		ownerInfo: owner{name: "Mert"},
	} // initializing struct with field names
	myEngine2.mpg = 40                                      // updating the value of the field
	fmt.Printf("myEngine2: %v, %T\n", myEngine2, myEngine2) // {30 10}, main.gasEngine

	// ----------- Slices and mehtods of a struct:
	var myEngine3 gasEngine = gasEngine{
		mpg:       30,
		gallons:   10,
		ownerInfo: owner{name: "Mert"},
	}
	fmt.Printf("myEngine3 miles left: %v\n", myEngine3.milesLeft()) // 300 (30 miles per gallon * 10 gallons)

	// ----------- Interfaces:
	var myElectricEngine electricEngine = electricEngine{
		mpk: 3,
		kwh: 10,
	}
	fmt.Printf("myElectricEngine miles left: %v\n", myElectricEngine.milesLeft()) // 30 (3 miles per kilowatt-hour * 10 kilowatt-hours)

	canMakeIt(myEngine3, 250)
	canMakeIt(myElectricEngine, 40)
}
