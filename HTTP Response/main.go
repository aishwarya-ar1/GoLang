//sending request to google and printing response

package main

import (
	"fmt"
	"net/http"
	"os"
)

//net package is for network interface
//package used here is 'http'
//error will only exists if one occurs
func main() {
	//always inclue URL with http
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
