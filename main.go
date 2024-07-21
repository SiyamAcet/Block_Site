package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

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
