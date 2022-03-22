package main

import (
	"crawler/crawler/frontend/controller"
	"net/http"
)

// main 执行
// go run crawler/starter.go
func main() {
	handler := controller.CreateSearchResultHandler("crawler/frontend/view/template.html")

	http.Handle("/search", handler)
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
