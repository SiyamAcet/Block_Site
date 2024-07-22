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
	view.Execute(w, nil)

}

func Test(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	check := r.FormValue("check")
	select_input := r.FormValue("select")

	fmt.Println(check, select_input)

}

func main() {
	r := httprouter.New()
	r.GET("/", MainPage)
	r.POST("/test", Test)
	http.ListenAndServe(":8080", r)
}
