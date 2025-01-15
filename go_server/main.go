package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 Not Found", http.StatusNotFound)
	}
	if req.Method != "GET" {
		http.Error(res, "Method is not Supported", http.StatusNotFound)
	}

	fmt.Fprintf(res, "Hello!")
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.Redirect(res, req, "/form.html", 301)
	} else if req.Method == "POST" {

		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm Err: %v", err)
			return
		}

		fmt.Fprintf(res, "POST Request Success!\n")
		name := req.FormValue("name")
		address := req.FormValue("address")
		fmt.Fprintf(res, "Name=%s\n", name)
		fmt.Fprintf(res, "Address=%s\n", address)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static/"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
