package structs

import "time"

type Fidelity struct{
	Id int `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Plan string `json:"plan"`
}

