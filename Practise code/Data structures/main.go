//struct decleration types

package main

import "fmt"

type person struct { //defining our struct
	firstName string
	lastName  string
}

func main() {
	// def struct method 1 :go assumes alex as first and anderson as second because of order
	//alex := person {"Alex", "Anderson"}
	// Method 2
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	//method 3
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	// printf kicks out struct and shows first and last name var
	fmt.Printf("%+v", alex)
}
