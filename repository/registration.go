package repository

import (
	"api.1/interfaces"
	"api.1/structs"
)

type RegistrationRepo struct {
	db interfaces.Repo
}

func (r RegistrationRepo) Insert(info structs.Registration) (structs.Registration, error) {
	var infoReturn structs.Registration

	err := r.db.QueryRow(`INSERT INTO Registration (id, name, cpf, "e-mail", telephone) VALUES ($1,$2,$3,$4,$5) RETURNING id,name,cpf,"e-mail",telephone`, []any{info.Id, info.Name, info.Cpf, info.Email, info.Telephone}, &infoReturn.Id, &infoReturn.Name, &infoReturn.Cpf, &infoReturn.Email, &infoReturn.Telephone)

	return infoReturn, err
}

func (r RegistrationRepo) Get(info int) (structs.Registration, error) {
	var infoReturn structs.Registration

	err := r.db.QueryRow(`SELECT * FROM Registration WHERE id=$1`, []any{info}, &infoReturn.Id, &infoReturn.Name, &infoReturn.Cpf, &infoReturn.Email, &infoReturn.Telephone)

	return infoReturn, err
}

func (r RegistrationRepo) Update(info structs.Registration) (structs.Registration, error) {
	var infoReturn structs.Registration

	err := r.db.QueryRow(`UPDATE Registration SET name=$2, cpf=$3, "e-mail"=$4, telephone=$5 WHERE id=$1 RETURNING id, name,cpf,"e-mail",telephone`, []any{info.Id, info.Name, info.Cpf, info.Email, info.Telephone}, &infoReturn.Id, &infoReturn.Name, &infoReturn.Cpf, &infoReturn.Email, &infoReturn.Telephone)

	return infoReturn, err

}

func (r RegistrationRepo) Delete(info structs.Registration) error {
	err := r.db.QueryRow(`DELETE FROM Registration WHERE id=$1`, []any{info.Id})

	return err
}

func NewRegistrationRepo(db interfaces.Repo) RegistrationRepo {
	return RegistrationRepo{db: db}
}
