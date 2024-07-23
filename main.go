package main

import (
	"fmt"
	"net/http"

	admin_models "example.com/my-blog/admin/models"
	"example.com/my-blog/config"
)

func main() {

	// admin_models.Post{}.Migrate()

	// admin_models.Post{
	// 	Title:       "Post 1",
	// 	Slug:        "post-1",
	// 	Description: "This is post 1",
	// 	Content:     "This is post 1 content",
	// 	Picture_url: "https://via.placeholder.com/150",
	// 	CategoryID:  1,
	// }.Add()

	// post := admin_models.Post{}.Get("description = ?", "This is post 1")

	// fmt.Println(post.Title)

	fmt.Println(admin_models.Post{}.GetAll())

	http.ListenAndServe(":8080", config.Routes())

}
