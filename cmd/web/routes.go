package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	sessionMiddleware := app.sessionManager.LoadAndSave

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /{$}",
		sessionMiddleware(http.HandlerFunc(app.home)),
	)

	mux.Handle("GET /snippet/view/{id}",
		sessionMiddleware(http.HandlerFunc(app.snippetView)),
	)

	mux.Handle("GET /snippet/create",
		sessionMiddleware(app.requireAuthentication(http.HandlerFunc(app.snippetCreate))),
	)

	mux.Handle("POST /snippet/create",
		sessionMiddleware(app.requireAuthentication(http.HandlerFunc(app.snippetCreatePost))),
	)

	mux.Handle("GET /user/signup", sessionMiddleware(http.HandlerFunc(app.userSignup)))
	mux.Handle("POST /user/signup", sessionMiddleware(http.HandlerFunc(app.userSignupPost)))
	mux.Handle("GET /user/login", sessionMiddleware(http.HandlerFunc(app.userLogin)))
	mux.Handle("POST /user/login", sessionMiddleware(http.HandlerFunc(app.userLoginPost)))
	mux.Handle("POST /user/logout", sessionMiddleware(http.HandlerFunc(app.userLogoutPost)))
	// Pass the servemux as the 'next' parameter to the commonHeaders middleware.
	// Because commonHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.
	return app.recoverPanic(app.logRequest(commonHeaders(mux)))
}
