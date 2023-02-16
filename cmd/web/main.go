// /greeting handler called greeting and the handler will output a message
//Welcome to my page

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"time"
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
		Addr:    *addr,
		Handler: app.routes(), //we created this routes server

	}

	log.Printf("Starting Server on port %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err) //should never be reached
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	//use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//lets ping the db
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
