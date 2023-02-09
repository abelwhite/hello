// /greeting handler called greeting and the handler will output a message
//Welcome to my page

package main

import (
	"flag"
	"log"
	"net/http"
)

// create a new type
type application struct{}

func main() {
	//Create a flag for specifing the port number when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	flag.Parse()

	//create a new instance of the application type
	app := &application{}



	//create a customized server
	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(), //we created this routes
		
	}

	log.Printf("Starting Server on port %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err) //should never be reached
}
