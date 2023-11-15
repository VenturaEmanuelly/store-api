package structs

type Registration struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Cpf string `json:"cpf"`
	Email string `json:"email"`
	Telephone string `json:"telephone"`
	Fidelity Fidelity `json:"fidelity,omitempty"`
}