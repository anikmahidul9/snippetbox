package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//Create a file server which serves files out of the "./ui/static" directory. The path given to the http.dir function is relative to the project root directory.

fileServer := http.FileServer(http.Dir("./ui/static"))

	//Use the mux.Handle() method to register the file server as the handler for all URL paths that start with "/static/". We use the http.StripPrefix() function to create a new handler which will remove the "/static" prefix from the request URL's path before passing the request to the file server. This is necessary because our static files are located in the "./ui/static" directory, but we want them to be accessible via URLs that start with "/static/".

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{$}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Server starting on port : 4000")

	err :=http.ListenAndServe(":4000",mux)
	if err !=nil{
		log.Fatal(err)
	}
}
