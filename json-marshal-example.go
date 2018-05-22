
package main

import(
	"log"
	"encoding/json"
	"strconv"

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

	// Ex: 1. struct into json (byte/escaped string)
	m, err := json.Marshal(animal1)
	if err != nil {log.Println(err)}
	log.Print("marshaled : ",string(m))

	// Ex: 2.  marshaled json (byte/escaped string) into struct
	err = json.Unmarshal(m, &animal2)
	log.Print("unmarshaled : ", animal2, " ",animal2.Name, " ", animal2.Order)
	
	// Ex: 3.  raw json (byte/escaped string) into struct
	b := []byte(`{"Name":"dallas","Order":"texas"}`)
	err = json.Unmarshal(b, &animal3)
	log.Print("unmarshaled : ", animal3, " ",animal3.Name, " ", animal3.Order)

	// Ex: 4.  scaped string into struct
	bs := []byte(`"{\"name\":\"texas\", \"order\":\"porkrinds\"}"`)
	s,_ := strconv.Unquote(string(bs))
	err = json.Unmarshal([]byte(s), &animal4)
	log.Println("escaped byte ",animal4)
	
	
	

}
