package entities

import (
	"fmt"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"division"`
	ProjectID  int    `json:"project"`
}

func (e Employee) String() string {
	return fmt.Sprintf("%s is working in %s department on project: %d.", e.Name, e.Department, e.ProjectID)
}
