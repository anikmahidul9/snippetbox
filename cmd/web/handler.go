package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snippetbox.anikmahidul9/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	//Use the template.ParseFiles() function to read  the template file into a template set http.Error() function to send an internal server error response to the user, and then return from the handler so no subsequent code is executed.
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Call the newTemplateData() helper to get a templateData struct containing
// the 'default' data (which for now is just the current year), and add the
// snippets slice to it.
data := app.newTemplateData(r)
data.Snippets = snippets
	
  app.render(w,r,http.StatusOK,"home.tmpl",data)
}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.logger.Error("Invalid snippet ID", "method", r.Method, "url", r.URL.RequestURI())
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	data := app.newTemplateData(r)
	data.Snippet = *snippet
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.logger.Error("Failed to insert snippet", "method", r.Method, "url", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
