
// github.com/auth0/go-jwt-middleware/blob

package main

import(
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni" // a fork?
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	StartServer()
}

func StartServer(){
	r := mux.NewRouter()

	jwtMiddleware :=  jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{},error){
			return []byte("secret"),nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	r.HandleFunc("/ping", PingHandler)
	r.Handle("/secured/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(SecuredPingHandler)),
	))
	http.Handle("/",r)
	http.ListenAndServe(":3001", nil)
}

type Response struct {
	Text string `json:"text"`
}

func respondJson(text string, w http.ResponseWriter){
	response := Response(text)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func PingHandler(w http.ResponseWriter, r *http.Request){
	respondJson("not authenticated response", w)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request){
	respondJson("authenticated response", w)
}
























