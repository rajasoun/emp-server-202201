package service

//go:generate mockgen -source=$GOFILE -package $GOPACKAGE -destination=mock_$GOPACKAGE.go

import "algogrit.com/httpex/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(newEmployee entities.Employee) (*entities.Employee, error)
}
