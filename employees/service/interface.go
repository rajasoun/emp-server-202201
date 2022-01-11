package service

import "algogrit.com/httpex/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(newEmployee entities.Employee) (*entities.Employee, error)
}
