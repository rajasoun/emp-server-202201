package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"algogrit.com/httpex/entities"
)

func (h EmployeeHandler) Index(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	employees, _ := h.empSvcV1.Index()

	json.NewEncoder(w).Encode(employees)
}

func (h EmployeeHandler) Create(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp, _ := h.empSvcV1.Create(newEmployee)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newEmp)

}
