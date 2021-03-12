//using channels to enable communication between go(child) routines and main routines.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// make (built-in func) creates a value of determined type
	c := make(chan string)
	//c channel is applicable only inside the main function
	for _, link := range links {
		//go keyword is always placed in front of function calls
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		//Output will not be printed until data is received from other go(child) routines. Hence, we add iteration = no. of links available.
	}
}

//checks if link is responding to traffic
// channel is abbreviated as chan or c
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

//link to course instructor's completed code
//https://github.com/StephenGrider/GoCasts/tree/master/code/channelss
