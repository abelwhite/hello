package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// create a multiplexer
	router := httprouter.New() //httprouter from juienschmidft/httprouter
	// create a file server
	// filer server is to
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //remove "resources"

	router.HandlerFunc(http.MethodGet, "/create", app.Greeting) //provide string and hander
	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodGet, "/about", app.About)
	router.HandlerFunc(http.MethodPost, "/create", app.MessageCreate)

	return router
}
