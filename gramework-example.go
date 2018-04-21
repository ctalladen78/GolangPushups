package main

import (
	"encoding/json"
	"fmt"

	"github.com/gramework/gramework"
	"github.com/valyala/fasthttp"
)

var (
	templates = "templates"
)

// https://gobyexample.com/json
// use jsoniter for fast json
type jsonResponse struct {
	Name       string
	Age        int
	Attributes []string
}

// https://github.com/gramework/gramework
func main() {
	app := gramework.New()

	// TODO serve html with vue.js from file system
	app.GET("/", app.ServeDir("templates"))
	// app.GET("/", app.ServeDirCustom("templates", 0, true, false, []string{"index.html", "about.html"}))

	// app.JSON
	app.JSON("/jsonval", func(fctx *fasthttp.RequestCtx) {

		// TODO encode json into string (marshall)
		jsonRes := &jsonResponse{
			Name: "Cyrus",
			Age:  40,
			Attributes: []string{
				"artistic",
				"musical",
				"kind",
			},
		}

		res, _ := json.Marshal(jsonRes)

		fctx.WriteString(string(res))
	})

	fmt.Println("listening on localhost:8000")
	app.ListenAndServe("localhost:8000")

}
