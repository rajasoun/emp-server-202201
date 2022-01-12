package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

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

	validate := validator.New()
	errs := validate.Struct(newEmployee)

	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, errs)
		return
	}

	newEmp, _ := h.empSvcV1.Create(newEmployee)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newEmp)
}
