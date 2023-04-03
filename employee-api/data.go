package main

var employees = []Employee{
	{ID: 1, Name: "Melih", Department: "Engineering"},
	{ID: 2, Name: "Mert", Department: "Engineering"},
	{ID: 3, Name: "Aleyna", Department: "Engineering"},
}

type Employee struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Department string `json:"department,omitempty"`
}

type CreateEmployeeRequest struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}
