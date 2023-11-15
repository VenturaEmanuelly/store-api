package repository

import (
	"time"

	"api.1/interfaces"
	"api.1/structs"
)

type FidelityRepo struct {
	db interfaces.Repo
}

func (f FidelityRepo) Upsert(fidelity structs.Fidelity) (structs.Fidelity, error) {
	var infoReturn structs.Fidelity

	err := f.db.QueryRow(`INSERT INTO fidelity (id, start_date, end_date, plan) VALUES ($1, $2, $3, $4 ) 
	ON CONFLICT (id) DO UPDATE SET start_date= $2, end_date= $3, plan= $4 RETURNING id, start_date, end_date, plan`, []any{fidelity.Id, fidelity.StartDate, fidelity.EndDate, fidelity.Plan}, &infoReturn.Id, &infoReturn.StartDate, &infoReturn.EndDate, &infoReturn.Plan)

	return infoReturn, err
}

func (f FidelityRepo) Get(info int) (structs.Fidelity, error) {
	var infoReturn structs.Fidelity

	err := f.db.QueryRow(`SELECT * FROM fidelity WHERE id=$1`, []any{info}, &infoReturn.Id, &infoReturn.StartDate, &infoReturn.EndDate, &infoReturn.Plan)

	return infoReturn, err
}

func (f FidelityRepo) UpdatePlan(endDate time.Time) error {
	err := f.db.QueryRow(`UPDATE fidelity SET plan = 'bronze' WHERE end_date <= $1 AND plan != 'bronze'`, []any{endDate})
	return err
}

func NewFidelityRepo(db interfaces.Repo) FidelityRepo {
	return FidelityRepo{db: db}
}
