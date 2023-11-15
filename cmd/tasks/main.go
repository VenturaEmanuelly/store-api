package main

import (
	driver "database/sql"
	"fmt"

	"api.1/repository"
	"api.1/tasks"
	_ "github.com/lib/pq"
)

func main() {
	dependecy, err := injectDependecy()
	if err != nil {
		return
	}

	err = dependecy.UpdatePlan()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func injectDependecy() (tasks.FidelityTask, error) {
	db, err := driver.Open("postgres", "host=localhost port=5432 user= store password=golang dbname= store sslmode=disable")
	if err != nil {
		return tasks.FidelityTask{}, err
	}
	
	repo := repository.NewRepository(db)
	fidelityRepo := repository.NewFidelityRepo(repo)
	return tasks.NewFidelityTask(fidelityRepo), nil
}
