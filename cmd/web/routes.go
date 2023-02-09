package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// create a multiplexer
	mux := http.NewServeMux() //multiplexer is to which handler func to call
	// create a file server
	// filer server is to
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //remove "resources"

	mux.HandleFunc("/greeting", app.Greeting) //provide string and hander
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/message/create", app.MessageCreate)
	
	return mux
}
