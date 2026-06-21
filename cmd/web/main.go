package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"snippetbox.anikmahidul9/internal/models"
)

type application struct {
	logger *slog.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	// Define new command line flag with the name "addr", a default value of ":4000" and a short description. The value of this flag will be stored in the addr variable.

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:123456@/snippetbox?parseTime=true", "SQL Data source")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error("Failed to create template cache", "error", err)
		os.Exit(1)
	}
	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	mux := http.NewServeMux()

	//Create a file server which serves files out of the "./ui/static" directory. The path given to the http.dir function is relative to the project root directory.

	fileServer := http.FileServer(http.Dir("./ui/static"))

	//Use the mux.Handle() method to register the file server as the handler for all URL paths that start with "/static/". We use the http.StripPrefix() function to create a new handler which will remove the "/static" prefix from the request URL's path before passing the request to the file server. This is necessary because our static files are located in the "./ui/static" directory, but we want them to be accessible via URLs that start with "/static/".

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So in this code, that means the addr variable
	// is actually a pointer, and we need to dereference it (i.e. prefix it with
	// the * symbol) before using it. Note that we're using the log.Printf()
	// function to interpolate the address with the log message.
	logger.Info("Starting server on", "addr", *addr)

	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

	
}
func openDB(dsn string) (*sql.DB, error) {
		db, err := sql.Open("mysql",dsn)
		if err != nil {
			return nil, err
		}
		if err = db.Ping(); err != nil {
			return nil, err
		}
		return db, nil
	}