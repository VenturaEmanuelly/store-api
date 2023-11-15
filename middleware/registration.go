package middleware

import (
	"errors"
	"strings"
	"unicode"

	"api.1/interfaces"
	"api.1/structs"
)

type RegistrationMiddleware struct {
	db interfaces.RegistrationRepo
	fidelityDB interfaces.FidelityRepo
}

func (r RegistrationMiddleware) Create(body structs.Registration) (structs.Registration, error) {

	err := r.telephoneChecker(body)
	if err != nil {
		return structs.Registration{}, err
	}

	r.nameChecker(&body)

	err = r.emailChecker(body)
	if err != nil {
		return structs.Registration{}, err
	}

	err = r.cpfChecker(body)
	if err != nil {
		return structs.Registration{}, err
	}

	return r.db.Insert(body)
}

func (r RegistrationMiddleware) telephoneChecker(info structs.Registration) error {
	telephoneToRune := []rune(info.Telephone)
	if string(telephoneToRune[0:2]) != "55" {
		return errors.New("the first 2 digits do not correspond to the Brazilian code")
	}
	return nil
}

func (r RegistrationMiddleware) nameChecker(info *structs.Registration) {
	nameToRune := []rune(info.Name)
	nameToRune[0] = unicode.ToUpper(nameToRune[0])
	info.Name = string(nameToRune)
}

func (r RegistrationMiddleware) emailChecker(info structs.Registration) error {
	if strings.Contains(info.Email, "@") && strings.Contains(info.Email, ".com") {
		return nil
	}
	return errors.New("Missing special characters: @, .com")
}

func (r RegistrationMiddleware) cpfChecker(info structs.Registration) error {
	cpfToRune := []rune(info.Cpf)
	if string(cpfToRune[3]) == "." && string(cpfToRune[7]) == "." && string(cpfToRune[11]) == "-" {
		return nil
	}
	return errors.New("You need to follow the character pattern in the CPF")
}

func (r RegistrationMiddleware) Read(body structs.Registration) (structs.Registration, error) {
	fidelityBody,_:= r.fidelityDB.Get(body.Id)

	registrationBody, err := r.db.Get(body.Id)
	if fidelityBody != (structs.Fidelity{}){
		registrationBody.Fidelity = fidelityBody
	}
	return registrationBody, err

}

func (r RegistrationMiddleware) Update(body structs.Registration) (structs.Registration, error) {
	return r.db.Update(body)
}

func (r RegistrationMiddleware) Delete(body structs.Registration) error {
	return r.db.Delete(body)
}

func NewRegistrationMiddleware(db interfaces.RegistrationRepo,fidelityDB interfaces.FidelityRepo) RegistrationMiddleware {
	return RegistrationMiddleware{db: db, fidelityDB: fidelityDB}
}
