//to print the body of the response from google
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//net package is for network interface
//package used here is 'http'
//error will only exists if one occurs

type logWriter struct{}

func main() {
	//always inclue URL with http
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//fmt.Println(resp)
	//using the read function to print the body
	// byte slice(type), no of elements and empty spaces should be defined with
	//read function reads data until byte slice is full
	//bs := make([]byte, 99999)
	//response struct implements the reader interface
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	//prints html that makes google.com
	//io.Copy(os.Stdout, resp.Body)
	// io.copy has a "read and write" function in-built.

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these many byte:", len(bs))
	return len(bs), nil

}
