package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"snippetbox.anikmahidul9/internal/models"
)

type application struct {
	logger         *slog.Logger
	users          *models.UserModel
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
	formDecoder    *form.Decoder
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

	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Cookie.Secure = true

	formDecoder := form.NewDecoder()

	app := &application{
		logger:         logger,
		users:          &models.UserModel{DB: db},
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  templateCache,
		sessionManager: sessionManager,
		formDecoder:    formDecoder,
	}

	// mux := http.NewServeMux()

	// //Create a file server which serves files out of the "./ui/static" directory. The path given to the http.dir function is relative to the project root directory.

	// fileServer := http.FileServer(http.Dir("./ui/static"))

	// //Use the mux.Handle() method to register the file server as the handler for all URL paths that start with "/static/". We use the http.StripPrefix() function to create a new handler which will remove the "/static" prefix from the request URL's path before passing the request to the file server. This is necessary because our static files are located in the "./ui/static" directory, but we want them to be accessible via URLs that start with "/static/".

	// mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// mux.HandleFunc("GET /{$}", app.home)
	// mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	// mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	// mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So in this code, that means the addr variable
	// is actually a pointer, and we need to dereference it (i.e. prefix it with
	// the * symbol) before using it. Note that we're using the log.Printf()
	// function to interpolate the address with the log message.

	// Initialize a new http.server struct, we set the Addr field to handler the server uses the same network address as before.

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	logger.Info("Starting server on", "addr", srv.Addr)

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	if err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
