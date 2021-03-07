//creating a structre and embedding another struct into it

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
	jim.print()
	jim.updateName("Jimmy")
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
	//OUTPUT 1
	//{firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:95000}}
	//didnt get updated - this is where pointers come into play
}

func (p person) print() {
	// prints person specified as "p" using '.print()'
	fmt.Printf("%+v", p)
}
