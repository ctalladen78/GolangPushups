
// chapter 1. Go Programming Language by Kernighan

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	url := "http://localhost:8000"
	
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprint(os.Stderr, "fetch: reading %s: %v \n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
