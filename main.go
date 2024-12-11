package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Photo  string 
}

type HeaderStruct struct {
	Links []string
}

func main() {
	headerLinks := []string{"home", "about", "contact"}

	posts, err := fetchPostsWithPhotos()
	if err != nil {
		fmt.Println("Error fetching posts with photos:", err)
		return
	}

	component := App(HeaderStruct{Links: headerLinks}, posts)

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

