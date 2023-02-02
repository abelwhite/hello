// /greeting handler called greeting and the handler will output a message
//Welcome to my page

package main

import (
	"log"
	"net/http"

	"github.com/abelwhite/hello/handlers"
)

func main() {
	//create a multiplexer
	mux := http.NewServeMux() //multiplexer is to which handler func to call
	//create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //remove "resources"

	mux.HandleFunc("/greeting", handlers.Greeting) //provide string and hander
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/message/create", handlers.MessageCreate)
	// /home
	// /about
	//

	log.Println("Starting Server on port:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err) //should never be reached
}
