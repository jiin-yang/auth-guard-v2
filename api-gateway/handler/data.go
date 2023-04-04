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

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
