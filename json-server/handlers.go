package main

// handler functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// explicit typing
type emailTemplate struct {
	personalizations string
	from             []string // map [string]interface{}
	reply_to         []string
}

// dynamic typing using empty interface
const jsonExample = `{ 'personalizations': 
	[ { 'to': [ { email: 'john.doe@example.com', 'name': 'John Doe' } ],
		'subject': 'Hello, World!' } ],
   'from': { 'email': 'sam.smith@example.com', 'name': 'Sam Smith' },
   'reply_to': { 'email': 'sam.smith@example.com', 'name': 'Sam Smith' } }`

type Animal struct {
	Name  string
	Order string
}

type Animals struct {
	Animal []Animal `json:"animal"`
}

// https://blog.golang.org/json-and-go
//https://golang.org/src/encoding/json/example_test.go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	type ColorGroup struct {
		id     int      `json: id`
		Name   string   `json: name`
		Colors []string `json: colors`
	}
	group := ColorGroup{
		id:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	var group1 ColorGroup
	marshaled, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	unmarshaled := json.Unmarshal(marshaled, &group1)

	os.Stdout.Write(marshaled)
	log.Print("unmarshalled ", unmarshaled)
	log.Print("marshalled ", string(marshaled))
	// Output:
	// {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
}

// boilerplate
// using json.NewDecoder & strings.NewReader
// https://attilaolah.eu/2013/11/29/json-decoding-in-go/
func jsonStreamHandler(w http.ResponseWriter, r *http.Request) {
	var msg interface{}

	const jsonStream = `{"Message": "Hello", "Array": [1, 2, 3]}`
	dec := json.NewDecoder(strings.NewReader(jsonStream)) // TODO do not use this on large datasets
	// log.Println("json example", dec)

	if err := dec.Decode(&msg); err != nil {
		return
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}
	w.Write(b)

}

// echoHandler reads a JSON object from the body, and writes it back out.
// https://gist.github.com/campoy/7b44f6ec2d9e82d956d34b4989b33192
func echoHandler(w http.ResponseWriter, r *http.Request) {
	var msg interface{}
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			// log.Fatal(w, http.StatusBadRequest, "Body was not valid JSON: %v", err)
			log.Fatal("Body was not valid JSON: %v", err)
			return
		}
		// log.Fatal(w, http.StatusInternalServerError, "Could not get body: %v", err)
		log.Fatal("Could not get body: %v", err)
		return
	}

	b, err := json.Marshal(msg)
	if err != nil {
		// log.Fatal(w, http.StatusInternalServerError, "Could not marshal JSON: %v", err)
		log.Fatal("Could not marshal JSON: %v", err)
		return
	}
	log.Println("json response: ", string(b))
	w.Write(b)
}

// errorf writes a swagger-compliant error response.
func errorHandler(w http.ResponseWriter, code int, format string, a ...interface{}) {
	var out struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	out.Code = code
	out.Message = fmt.Sprintf(format, a...)

	b, err := json.Marshal(out)
	if err != nil {
		http.Error(w, `{"code": 500, "message": "Could not format JSON for original message."}`, 500)
		return
	}

	http.Error(w, string(b), code)
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.sendgrid.com/v3/mail/send"

	// TODO how does strings.NewReader parse strings into json
	payload := strings.NewReader("{\"personalizations\":[{\"to\":[{\"email\":\"cyrus@tasselvr.com\",\"name\":\"John Doe\"}],\"subject\":\"Hello, World!\"}],\"from\":{\"email\":\"sam.smith@example.com\",\"name\":\"Sam Smith\"},\"reply_to\":{\"email\":\"sam.smith@example.com\",\"name\":\"Sam Smith\"}}")

	req, _ := http.NewRequest("POST", url, payload)
	auth := "Bearer " + os.Getenv("SENDGRID_API_KEY")
	log.Println("auth ", auth)
	req.Header.Add("authorization", auth)
	req.Header.Add("content-type", "application/json")

	// res, _ := http.DefaultClient.Do(req)
	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// log.Println("request: ", r)
	ctx := appengine.NewContext(r)
	// ctx := r.Context()
	client := urlfetch.Client(ctx)

	resp, err := client.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	// TODO the post request to sendgrid is returning a 405 response due to malformed json payload

	fmt.Fprint(w, "SendGrid response: ", resp.Status)

	fmt.Println(resp)
	fmt.Println(string(body))
	defer resp.Body.Close()
}
