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

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	var empRepo = repository.NewInMem()
	var empSvcV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvcV1)

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port:", 8000, "...")

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
