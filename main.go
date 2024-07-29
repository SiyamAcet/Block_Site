package main

import (
	"net/http"

	"example.com/my-blog/config"

	admin_models "example.com/my-blog/admin/models"
)

func main() {

	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()

	http.ListenAndServe(":8080", config.Routes())

}
