package repository

import (
	"database/sql"
	"log"

	"algogrit.com/httpex/entities"
	_ "github.com/lib/pq"
)

type sqlRepo struct {
	db *sql.DB
}

func (repo *sqlRepo) FindAll() ([]entities.Employee, error) {
	rows, err := repo.db.Query("SELECT * FROM employees")

	if err != nil {
		return nil, err
	}

	employees := []entities.Employee{}

	for rows.Next() {
		var employee entities.Employee
		rows.Scan(&employee.ID, &employee.Name, &employee.Department, &employee.ProjectID)

		employees = append(employees, employee)
	}

	return employees, nil
}

func (repo *sqlRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	rows, err := repo.db.Query("SELECT count(*) FROM employees")

	if err != nil {
		return nil, err
	}

	rows.Next()

	var count int
	rows.Scan(&count)

	newEmp.ID = count + 1

	_, err = repo.db.Exec("INSERT INTO employees (id, name, department, project_id) VALUES ($1, $2, $3, $4)", newEmp.ID, newEmp.Name, newEmp.Department, newEmp.ProjectID)

	if err != nil {
		return nil, err
	}

	return &newEmp, nil
}

func NewSql() EmployeeRepository {
	db, err := sql.Open("postgres", "postgres://localhost:5432/template1?sslmode=disable")

	if err != nil {
		log.Fatalln("Unable to connect to DB:", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS employees (id numeric primary key, name text, department text, project_id numeric)")

	if err != nil {
		log.Fatalln("Unable to create table:", err)
	}

	return &sqlRepo{db}
}
