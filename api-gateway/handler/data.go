package handler

type Employee struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Department string `json:"department,omitempty"`
}

type CreateEmployeeRequest struct {
	Name       string `json:"name,omitempty"`
	Department string `json:"department,omitempty"`
}

type CreateEmployeeResponse struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Department string `json:"department,omitempty"`
}
