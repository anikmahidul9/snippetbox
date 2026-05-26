package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	//Use the template.ParseFiles() function to read  the template file into a template set http.Error() function to send an internal server error response to the user, and then return from the handler so no subsequent code is executed.
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil{
		log.Println(err.Error())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
		return
	}
	//Then we use the Execute() method on the template set to write the template content as the response body. the last parameter to execute is the dynamic data that we want to pass in to the template. In this case we don't have any dynamic data to pass in, so we use nil.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil{
		log.Println(err.Error())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}
