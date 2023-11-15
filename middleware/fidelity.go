package middleware

import (
	"errors"
	"time"

	"api.1/interfaces"
	"api.1/structs"
)

type FidelityMiddleware struct {
	db interfaces.FidelityRepo
}

func (f FidelityMiddleware) Create(body structs.Fidelity) (structs.Fidelity, error) {
	err := f.checkFidelity(body)
		if err != nil{
			return structs.Fidelity{}, err 
		}
	return f.db.Upsert(body)

}

func (f FidelityMiddleware) checkFidelity(body structs.Fidelity) (error) {
	if body.Id == 0{
		return errors.New("missing id")
	}
	
	if body.StartDate == (time.Time{}) {
		return errors.New("missing startDate")
	}

	if body.EndDate == (time.Time{}) {
		return errors.New("missing endDate")
	}

	if body.Plan == "" {
		return errors.New("missing plan")
	}

	return nil
}

func (f FidelityMiddleware) Read(body structs.Fidelity) (structs.Fidelity, error) {
	return f.db.Get(body.Id)
}

func NewFidelityMiddleware(db interfaces.FidelityRepo) FidelityMiddleware{
	return FidelityMiddleware{db: db}
}