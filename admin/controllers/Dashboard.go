package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"example.com/my-blog/admin/helpers"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)

}

func (dashboard Dashboard) NewItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)

}
