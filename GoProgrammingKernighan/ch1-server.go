
// implement http server

package main

import (
	"net/http"
	"fmt"
	"log"
)

// in case the localhost port doesnt close
//http://stackoverflow.com/questions/3855127/find-and-kill-process-locking-port-3000-on-mac
func main(){
	http.HandleFunc("/", handlerFunc)
	fmt.Println("running server")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
}
