package interfaces

import (
	"time"

	"api.1/structs"
)

type RegistrationRepo interface {
	Insert(structs.Registration) (structs.Registration, error)
	Get(int) (structs.Registration, error)
	Update(structs.Registration) (structs.Registration, error)
	Delete(structs.Registration) error
}

type FidelityRepo interface {
	Upsert(structs.Fidelity) (structs.Fidelity, error)
	Get(int) (structs.Fidelity, error)
	UpdatePlan(time.Time) error
}

type Repo interface {
	QueryRow(query string, args []any, dest ...any) error
}

type RegistrationMiddleware interface {
	Create(structs.Registration) (structs.Registration, error)
	Read(structs.Registration) (structs.Registration, error)
	Update(structs.Registration) (structs.Registration, error)
	Delete(structs.Registration) error
}

type FidelityMiddleware interface {
	Create(structs.Fidelity) (structs.Fidelity, error)
	Read(structs.Fidelity) (structs.Fidelity, error)
}