package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"example.com/my-blog/admin/helpers"
	"github.com/julienschmidt/httprouter"
)

type Userops struct{}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("userops/login")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)

}

func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username, password)

}
