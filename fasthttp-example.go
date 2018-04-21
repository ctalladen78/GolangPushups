package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/test/testvar", Subpage)
	log.Println("Server started listening on port 8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))

	// https://github.com/valyala/fasthttp/blob/master/examples/fileserver/fileserver.go
	// use a goroutine to run server in a separate thread
	/*
	   go func(){
	   	log.Printf("starting server on %q", *addrTLS)
	   	go func(){
	   		if err := fasthttp.ListenAndServeTLS(*addrTLS, *certFile, *keyFile, requestHandler){
	   		log.Fatalf("error in fasthttp.ListenAndServeTLS", err)
	   }
	   	}
	   }
	*/
}

// Index show index page
// http://github.com/valyala/fasthttp#switching-from-nethttp-to-fasthttp
// https://github.com/valyala/fasthttp/blob/master/examples/fileserver/fileserver.go
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome! test\n")
	/*
		const tpl = `
		<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>{{.Title}}</title>
			</head>
			<body>
				{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
			</body>
		</html>`
		t, err := template.New("webpage").Parse(tpl)
	*/

}

// Subpage show test params
func Subpage(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "testing params, %s \n", ctx.UserValue("name"))

}
