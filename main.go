package main

import (
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func MainPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)

}

func Test(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(10 << 20)
	file, header, _ := r.FormFile("file")
	f, _ := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
}

func main() {
	r := httprouter.New()
	r.GET("/", MainPage)
	r.POST("/test", Test)
	http.ListenAndServe(":8080", r)
}
