package http

import (
	"algogrit.com/httpex/employees/service"

	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	empSvcV1    service.EmployeeService
	*mux.Router // Embedding => Inheritance
}

func (e *EmployeeHandler) SetupRoutes(r *mux.Router) {
	e.Router = r

	r.HandleFunc("/v1/employees", e.Index).Methods("GET")
	r.HandleFunc("/v1/employees", e.Create).Methods("POST")
}

func NewHandler(empSvcV1 service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{empSvcV1: empSvcV1}

	h.SetupRoutes(mux.NewRouter())

	return h
}
