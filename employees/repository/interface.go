package repository

//go:generate mockgen -source=$GOFILE -package $GOPACKAGE -destination=mock_$GOPACKAGE.go

import "algogrit.com/httpex/entities"

type EmployeeRepository interface {
	FindAll() ([]entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
