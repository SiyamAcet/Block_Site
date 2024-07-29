package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"example.com/my-blog/admin/helpers"
	"example.com/my-blog/admin/models"
	"github.com/julienschmidt/httprouter"
)

type Categories struct{}

func (categories Categories) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("categories/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)

	view.ExecuteTemplate(w, "index", data)
}
