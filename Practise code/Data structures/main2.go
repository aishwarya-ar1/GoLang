//creating a structre and embedding another struct into it
//after inserting POINTERS
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct { //defining our struct
	firstName string
	lastName  string
	contact   contactInfo
	//method 2 of embedding struct
	//contactInfo without "contact"
}

//while declaring multi line structs, every line must contain a comma - even after } and last line (command)
func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		//method 2
		//contactInfo:contactInfo
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 95000,
		},
	}
	//jimPointer := &jim
	//shortcut to pointers
	//from jimpointer to jim itself in both methods
	jim.updateName("Jimmy")
	jim.print()
}

//pointerToPerson becomes a value and *person is the type of pointer pointing at it
//here,jimpointer and pointertoperson are same (value)
func (pointerToPerson *person) updateName(newFirstName string) {
	//pointertoperson becomes a value as * is used here
	(*pointerToPerson).firstName = newFirstName
	//OUTPUT 1
	//{firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:95000}}
	//didnt get updated - this is where pointers come into play
}

func (p person) print() {
	// prints person specified as "p" using '.print()'
	fmt.Printf("%+v", p)
}
