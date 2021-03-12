//better way to check the links - parallel approach
//GO routines and channels come into play while checking sending requests to all links at once as soon as code is run!
package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		//go keyword is always placed in front of function calls
		go checkLink(link)
	}
}

//checks if link is responding to traffic
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down! Check back later!")
		return
	}

	fmt.Println(link, " is up!!")
}

// does not print an output like the other but compiles successfully
//error free code - without channel
