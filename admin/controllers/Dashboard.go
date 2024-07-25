package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gosimple/slug"

	"example.com/my-blog/admin/helpers"
	"example.com/my-blog/admin/models"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()

	view.ExecuteTemplate(w, "index", data)

}

func (dashboard Dashboard) NewItems(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)

}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	category_id, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	//upload
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test1")

	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	fmt.Println("test2")

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		fmt.Println(err)
		return
	}

	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  category_id,
		Content:     content,
		Picture_url: "uploads/" + header.Filename,
	}.Add()

	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post := models.Post{}.Get((params.ByName("id")))
	post.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}

func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(params.ByName("id"))

	view.ExecuteTemplate(w, "index", data)

}
