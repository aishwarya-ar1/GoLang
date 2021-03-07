package main

import "fmt"

func main() {
	// declaring with [key]values are string type
	//method 1
	colors := map[string]string{
		"red":   "#ff000",
		"green": "4bf745",
		"white": "#11223g",
	} //does not require , before bracket like struct

	printMap(colors) //copy of ref to map ; not method
	//method 2
	//var colors map [string]string

	//method 3 - uses built in functions
	// := make(map[string]string)
	//colors["white"] = "#fffff"
	//colors["red"] = "#12345f"
	//fmt.Println(colors)

	//delete function
	//delete(colors, "red")
	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
