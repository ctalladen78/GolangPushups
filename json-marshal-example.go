
package main

import(
	"log"
	"encoding/json"
	"strconv"
	"strings"
	"net/http"
	"fmt"
	"bytes"

)

func main(){
	type Animal struct {
		Name string `json:"name"`
		Order string `json"order"`
	}

	animal1 := Animal{Name: "austin", Order: "texas"}
	var animal2 Animal
	var animal3 Animal
	var animal4 Animal
	//var animal5 Animal
	//var animal6 Animal
	log.Println("animal 1: ", animal1)

	// Ex: 1. marshal struct into json (byte/escaped string)
	// marshal returns []byte slice
	m, err := json.Marshal(animal1)
	if err != nil {log.Println(err)}
	log.Print("marshaled : ",string(m))

	// Ex: 2.  unmarshaled json (byte/escaped string) into struct
	// unmarshal automagically instantiates correctly formed []byte slice into a similarly formed struct
	err = json.Unmarshal(m, &animal2)
	log.Print("unmarshaled : ", animal2, " ",animal2.Name, " ", animal2.Order)
	
	// Ex: 3.  unmarshal raw json (byte/escaped string) into struct
	b := []byte(`{"Name":"dallas","Order":"texas"}`)
	err = json.Unmarshal(b, &animal3)
	log.Print("unmarshaled : ", animal3, " ",animal3.Name, " ", animal3.Order)

	// Ex: 4.  unmarshal escaped string into struct
	bs := []byte(`"{\"name\":\"texas\", \"order\":\"porkrinds\"}"`)
	s,_ := strconv.Unquote(string(bs))
	err = json.Unmarshal([]byte(s), &animal4)
	log.Println("escaped byte ",animal4)
	
	// Ex. 5. using json.NewDecoder to conver string literal into empty interface
	// we need to read a string literal 
	var msg interface{}
	var stream = `{"city": "austin","state":"texas"}`
	decoded := json.NewDecoder(strings.NewReader(stream))
	if err := decoded.Decode(&msg); err != nil {
		log.Fatal(err)
	}
	// in order to access the value inside msg we have to do 
	// some extra parsing which is cumbersome compared to javascript
	log.Println("Ex.5 ", msg)
	

	// Ex. 6. using json.NewDecoder from incoming remote json request 
	// we can use curl for testing this function
	// $ curl localhost:8080 -d {"name":"san antonio", "order":"texas"}
	// here we can access a simple json
	/*
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.Body == nil {
			http.Error(w, "pls send a well formed json request body", http.StatusNotAcceptable)
return
		}
		err := json.NewDecoder(r.Body).Decode(&animal5)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		fmt.Println("Ex.6 ",animal5.Name)

	})
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
	
	// Ex. 7 lets make a client in memory that will
	// make a http request with a json body
	ex7 := Animal{Name:"el paso", Order:"texas"}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(ex7) // we encode a  byte buffer
	response, err := http.Post("http://localhost:8080","test", buf)
	if err != nil { fmt.Println(err.Error())}
	fmt.Println("Ex. 7", response)

	// Ex. 8 using github.com/tidwall/gjson
	// we can do some advanced querying for complex json data 
	
}

