package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// MainPage is a handler function for the main page of the website.
// It takes in the http.ResponseWriter, http.Request, and httprouter.Params as parameters.
// It parses the "index.html" template file and executes it with the provided data.
// The data is retrieved from the "slug" parameter in the URL.
// Finally, it prints "MainPage" to the console.
func MainPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, _ := template.ParseFiles("index.html")
	data := params.ByName("slug")
	view.Execute(w, data)
	fmt.Println("MainPage")

}

func main() {
	r := httprouter.New()
	r.GET("/posts/:slug", MainPage)
	http.ListenAndServe(":8080", r)
}
