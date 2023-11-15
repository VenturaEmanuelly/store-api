package main

import (
	driver "database/sql"

	"api.1/handlers"
	"api.1/middleware"
	"api.1/repository"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	registration, fidelity,  err := injectDependecy()
	if err != nil {
		return
	}

	registrationHandler := handlers.NewRegistrationHandler(registration)
	fidelityHandler := handlers.NewFidelidadeHandler(fidelity)

	app := fiber.New()
	app.Post("/", registrationHandler.HandlePost)
	app.Get("/", registrationHandler.HandleGet)
	app.Put("/", registrationHandler.HandlePut)
	app.Delete("/", registrationHandler.HandleDelete)

	app.Post("/fidelity", fidelityHandler.HandlePost)
	app.Get("/fidelity", fidelityHandler.HandleGet)

	app.Listen(":8080")

}

func injectDependecy() (middleware.RegistrationMiddleware, middleware.FidelityMiddleware, error) {
	db, err := driver.Open("postgres", "host=localhost port=5432 user= store password=golang dbname=store sslmode=disable")
	if err != nil {
		return middleware.RegistrationMiddleware{}, middleware.FidelityMiddleware{}, err
	}

	repo := repository.NewRepository(db)
	registrationSql := repository.NewRegistrationRepo(repo)
	fidelitySql := repository.NewFidelityRepo(repo)
	return middleware.NewRegistrationMiddleware(registrationSql, fidelitySql), middleware.NewFidelityMiddleware(fidelitySql), err

}
