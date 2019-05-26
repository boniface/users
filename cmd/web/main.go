package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"users/cmd/config"
	"users/cmd/web/controllers"
)

func main() {

	// Initialize a new template cache...
	templateCache, err := config.NewTemplateCache("./ui/html/")
	var path = "./cmd/web/views/html"
	if err != nil {
		fmt.Println(" There is a Error ", err)
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	env := &config.Env{
		ErrorLog:      log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
		InfoLog:       log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
		TemplateCache: templateCache,
		Path:          path,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: env.ErrorLog,
		Handler:  controllers.Controllers(env),
	}

	env.InfoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	error := srv.ListenAndServe()
	env.ErrorLog.Fatal(error)

}
