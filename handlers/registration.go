package handlers

import (
	"api.1/interfaces"
	"api.1/middleware"
	"api.1/structs"
	"github.com/gofiber/fiber/v2"
)

type RegistrationHandler struct {
	m interfaces.RegistrationMiddleware
}

func (r RegistrationHandler) HandlePost(ctx *fiber.Ctx) error {
	var info structs.Registration
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	registration, err := r.m.Create(info)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(registration)

}

func (r RegistrationHandler) HandleGet(ctx *fiber.Ctx) error {
	var info structs.Registration
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	registration, err := r.m.Read(info)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(registration)

}

func (r RegistrationHandler) HandlePut(ctx *fiber.Ctx) error {
	var info structs.Registration
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	registration, err := r.m.Update(info)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(registration)

}

func (r RegistrationHandler) HandleDelete(ctx *fiber.Ctx) error {
	var info structs.Registration
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	err = r.m.Delete(info)
	if err != nil {
		return err
	}
	return ctx.SendStatus(200)

}

func NewRegistrationHandler(m middleware.RegistrationMiddleware) RegistrationHandler {
	return RegistrationHandler{m: m}
}
