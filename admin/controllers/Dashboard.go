package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles("admin/views/index.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	view.Execute(w, nil)

}
