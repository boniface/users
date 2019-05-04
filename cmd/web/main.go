package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"users/cmd/config"
	"users/cmd/web/controllers"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	env := &config.Env{
		ErrorLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		InfoLog:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: env.ErrorLog,
		Handler:  controllers.Controllers(env),
	}

	env.InfoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	err := srv.ListenAndServe()
	env.ErrorLog.Fatal(err)

}
