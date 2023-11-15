package middleware

import (
	"testing"

	"api.1/structs"
)

func TestTelephoneCheck(t *testing.T) {
	create := RegistrationMiddleware{}
	t.Run("test telephone checker", func(t *testing.T) {
		if err := create.telephoneChecker(structs.Registration{Telephone: "55119785457"}); err != nil {
			t.Error(err)
		}
	})
	t.Run("test telephone checker with err", func(t *testing.T) {
		if err := create.telephoneChecker(structs.Registration{Telephone: "119785457"}); err == nil {
			t.Error(err)
		}
	})

}

func TestNameChecker(t *testing.T) {
	create :=  RegistrationMiddleware{}
	registration := structs.Registration{Name: "vincent"}
	t.Run("test name checker", func(t *testing.T) {
		create.nameChecker(&registration)
		if registration.Name != "Vincent" {
			t.Error("Error in the right name")
		}
	})

}

func TestEmailChecker(t *testing.T) {
	create := RegistrationMiddleware{}
	t.Run("test email checker", func(t *testing.T) {
		if err := create.emailChecker(structs.Registration{Email: "vincent78@.com"}); err != nil {
			t.Error(err)
		}
	})

	t.Run("test email checker with err", func(t *testing.T) {
		if err := create.emailChecker(structs.Registration{Email: "emanuellyv"}); err == nil {
			t.Error(err)
		}
	})
}

func TestCpfChecker(t *testing.T) {
	create :=  RegistrationMiddleware{}
	t.Run("test cpf checker", func(t *testing.T) {
		if err := create.cpfChecker(structs.Registration{Cpf: "412.412.412-4"}); err != nil {
			t.Error(err)
		}
	})

	t.Run("test cpf checker with err", func(t *testing.T) {
		if err := create.cpfChecker(structs.Registration{Cpf: "4124124124"}); err == nil {
			t.Error(err)
		}
	})
}
