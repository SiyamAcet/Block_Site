package main

import (
	"net/http"

	"example.com/my-blog/config"
)

func main() {

	// model view controller

	http.ListenAndServe(":8080", config.Routes())

}
