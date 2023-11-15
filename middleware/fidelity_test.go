package middleware

import (
	"testing"
	"time"

	"api.1/structs"
)

func TestCheckFidelity(t *testing.T) {
	create := FidelityMiddleware{}
	t.Run("test id checker", func(t *testing.T) {
		if err := create.checkFidelity(structs.Fidelity {Id: 0}); err == nil {
			t.Error(err)
		}
	})

	t.Run("test startDate checker", func(t *testing.T) {
		if err := create.checkFidelity(structs.Fidelity {Id: 8484, StartDate: time.Time{}}); err == nil {
			t.Error(err)
		}
	})

	t.Run("test endDate checker", func(t *testing.T) {
		if err := create.checkFidelity(structs.Fidelity {Id: 8484, StartDate: time.Now(), EndDate: time.Time{}}); err == nil {
			t.Error(err)
		}
	})

	t.Run("test plan checker", func(t *testing.T) {
		if err := create.checkFidelity(structs.Fidelity {Id: 8484, StartDate: time.Now(), EndDate: time.Now(), Plan: ""}); err == nil {
			t.Error(err)
		}
	})

	t.Run("test checker success", func(t *testing.T) {
		if err := create.checkFidelity(structs.Fidelity {Id: 8484, StartDate: time.Now(), EndDate: time.Now(), Plan: "gold"}); err != nil {
			t.Error(err)
		}
	})
}
