//File Name: handlers.go
package main

import (
	"net/http"

	"github.com/abelwhite/hello/helpers"
)

// creating handler function called greeting
// handler is called when we hit an end point
func (app *application) Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my page"))
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	//resttrict the end point
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	helpers.RenderTemplates(w, "./static/html/home.page.tmpl")
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {

}

func (app *application) MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//set header
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
