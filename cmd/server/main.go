package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	empHTTP "algogrit.com/httpex/employees/http"
	"algogrit.com/httpex/employees/repository"
	"algogrit.com/httpex/employees/service"
)

var port string

// var port *int

func init() {
	// port = flag.Int("port", 8000, "Port number on which is started")

	// flag.Parse()
	var ok bool
	port, ok = os.LookupEnv("PORT")

	if !ok {
		port = "8000"
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	// var empRepo = repository.NewInMem()
	var empRepo = repository.NewSql()
	var empSvcV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvcV1)

	empHandler.SetupRoutes(r)

	// log.Println("Starting server on port:", *port, "...")
	log.Println("Starting server on port:", port, "...")

	// err := http.ListenAndServe(":"+strconv.Itoa(*port), handlers.LoggingHandler(os.Stdout, r))
	err := http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))

	log.Fatalln("Unable to start server:", err)
}
