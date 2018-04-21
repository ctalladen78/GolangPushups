package main

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/iris"
)

type jsonRes struct {
	Name       string
	Attributes []string
}

func main() {
	app := iris.New()

	jsonRes := &jsonRes{
		Name: "Cyrus",
		Attributes: []string{
			"red",
			"yellow",
		},
	}
	jres, _ := json.Marshal(jsonRes)
	// files := ioutil.ReadDir("templates")

	fmt.Println("directory contents: ", jres)

	// to serve all files in the directory otherwise it only serves index.html
	// as an alternative of SPA you can take a look at the /routing/dynamic-path/root-wildcard
	assetHandler := app.StaticHandler("./templates", false, false)
	app.SPA(assetHandler)

	// or just serve index.html as it is:
	// app.Get("/{f:path}", func(ctx iris.Context) {
	// 	ctx.ServeFile("index.html", false)
	// })

	// register iris view engine
	app.RegisterView(iris.HTML("./templates", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./views/hello.html
		ctx.View("index.html")
	})

	app.Run(iris.Addr(":8000"))
}
