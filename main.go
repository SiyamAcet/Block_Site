package main

import (
	"net/http"
	"text/template"
)

// type Data struct {
// 	Message string
// 	Numbers []int
// }

// mainPage is the handler function for the main page.
// It parses the "index.html" template and executes it with no data.
// It also prints a message to the console when someone visits the main page.
func mainPage(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html", "navbar.html")

	view.ExecuteTemplate(w, "MainPage", nil)
}

func DetailPage(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("detail.html", "navbar.html")
	view.ExecuteTemplate(w, "DetailPage", nil)

}

func main() {
	// Register the mainPage function as the handler for the root ("/") route.
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/detail", DetailPage)

	// Start the HTTP server and listen on port 8080.
	http.ListenAndServe(":8080", nil)
}
