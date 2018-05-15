
// medium.com/@baijum/api-end-points...

package main

import (
	"net/http"
	jwtmiddleware "github.com/auth9/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main(){
	r1 := mux.NewRouter()
	r2 := mux.NewRouter()

	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"),nil

		},
		SigningMethod: jwt.SigningMethodHS256,

	})

	r1.HandleFunc("/api/with-auth", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("auth required \n"))

	}).Methods("GET")

	nn := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(r2))
	r1.PathPrefix("/api").Handler(nn)

	n := negroni.Classic()
	n.UseHandler(r1)
	n.Run(":8000")

}
