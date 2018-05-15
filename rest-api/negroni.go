
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

// GET ~/api/
func profileHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<p>This is content </p>")
	log.Println("accessed profile")
}

// GET ~/api/:var
func helloByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	fmt.Fprintf(w, "<p>This is %s </p>", ps.ByName("var"))
	log.Println("accessed with params")
}

// POST ~/api/
func postHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w, "<p>This no params </p>")
	log.Println("accessed with params")
}

// middleware function
func myMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	log.Println("auth middleware")	
	
	// call next handler
	next(rw, r)
	
}

func getUrlParams(router *httprouter.Router, req *http.Request) httprouter.Params {
	_, params, _ := router.Lookup(req.Method, req.URL.Path)
	return params
}

func callWithParams(rt *httprouter.Router, h func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		params := getUrlParams(rt, r)
		h(w, r, params)
	}
}


func main(){
	r := httprouter.New()
	r.POST("/login", postHandler)

	nh := negroni.New()
	nh.Use(n.HandlerFunc(myMiddleware))
	nh.UseHandlerFunc(profileHandler)
	r.Handler("GET", "/", nh)

	np := negroni.New()
	np.Use(negroni.HandlerFunc(myMiddleware))
	np.UseHandlerFunc(callWithParams(r, helloByName))
	r.Handler("GET", "/hello/:name", np)

	n := negroni.Classic()
	n.UseHandler(r)

	log.Fatal(http.ListenAndServe(":8080",n)


}







