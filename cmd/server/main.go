package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"algogrit.com/httpex/entities"
)

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "division": "%s", "project": %d}`, e.Name, e.Department, e.ProjectID)

// 	return []byte(jsonString), nil
// }

var employees = []entities.Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Pranav", "Cloud", 1002},
	{3, "Raja", "SRE", 10001},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// fmt.Fprintln(w, employees)

	// data, _ := json.Marshal(employees)
	// w.Write(data)

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newEmployee)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else {
// 		EmployeesIndexHandler(w, req)
// 	}
// }

// func LoggingMiddleware(h http.Handler) http.Handler {
// 	handlerFunc := func(w http.ResponseWriter, req *http.Request) {
// 		beginTime := time.Now()

// 		if /* request is acceptable */ {
// 			h.ServeHTTP(w, req)
// 		} else {
// 			w.WriteHeader(http.StatusNotAcceptable)
// 			w.Write()
// 		}

// 		log.Println(req.Method, req.URL, "Duration:", time.Since(beginTime))
// 	}

// 	return http.HandlerFunc(handlerFunc)
// }

func main() {
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
		// w.Write([]byte(msg))
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	// r.HandleFunc("/employees", EmployeesHandler)

	log.Println("Starting server on port:", 8000, "...")
	// http.ListenAndServe("127.0.0.1:8000", LoggingMiddleware(r))
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
