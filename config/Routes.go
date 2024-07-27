package config

import (
	"net/http"

	admin "example.com/my-blog/admin/controllers"
	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	// admin
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/add-new", admin.Dashboard{}.NewItems)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	// userops
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/login", admin.Userops{}.Login)
	//serve static files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))

	return r

}
