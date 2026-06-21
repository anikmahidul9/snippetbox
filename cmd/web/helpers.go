package main

import (
	"bytes"
	"fmt" // New import
	"net/http"
	"time"
)

// Create an newTemplateData() helper, which returns a templateData struct
// initialized with the current year. Note that we're not using the *http.Request
// parameter here at the moment, but we will do later in the book.
func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	// Retrieve the appropriate template set from the cache based on the page
	// name (like 'home.tmpl'). If no entry exists in the cache with the
	// provided name, then create a new error and call the serverError() helper
	// method that we made earlier and return.
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	// Write out the provided HTTP status code ('200 OK', '400 Bad Request' etc).
	w.WriteHeader(status)
	// Execute the template set and write the response body. Again, if there
	// is any error we call the serverError() helper.
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "url", r.URL.RequestURI())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}
