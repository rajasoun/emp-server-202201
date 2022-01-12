package entities

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name" validate:"required"`
	Department string `json:"division" validate:"required"`
	ProjectID  int    `json:"project"`
}

// func (e Employee) String() string {
// 	return fmt.Sprintf("%s is working in %s department on project: %d.", e.Name, e.Department, e.ProjectID)
// }
