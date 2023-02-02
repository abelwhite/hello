package handlers

import (
	"fmt"
	"net/http"
	"time"
)

// creating handler function called greeting
// handler is called when we hit an end point
func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my page"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	//resttrict the end point
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("This is my Home page"))
}

func About(w http.ResponseWriter, r *http.Request) {
	day := time.Now().Weekday()
	w.Write([]byte(fmt.Sprintf("Have a good %s.", day)))
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//set header
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Message"))
}
