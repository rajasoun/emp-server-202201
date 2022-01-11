package repository

import "algogrit.com/httpex/entities"

type EmployeeRepository interface {
	FindAll() ([]entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
