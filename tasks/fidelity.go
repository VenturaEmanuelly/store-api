package tasks

import (
	"time"

	"api.1/interfaces"
)

type FidelityTask struct {
	db interfaces.FidelityRepo
}

func (f FidelityTask) UpdatePlan() error {
	return f.db.UpdatePlan(time.Now())
}

func NewFidelityTask(db interfaces.FidelityRepo) FidelityTask {
	return FidelityTask{db: db}
}

