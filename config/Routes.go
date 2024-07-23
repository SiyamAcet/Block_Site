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
	//serve static files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))

	return r

}
