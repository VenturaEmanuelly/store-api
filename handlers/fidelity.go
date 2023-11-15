package handlers

import (
	"api.1/interfaces"
	"api.1/middleware"
	"api.1/structs"
	"github.com/gofiber/fiber/v2"
)

type FidelityHandler struct {
	m interfaces.FidelityMiddleware
}

func (f FidelityHandler) HandlePost(ctx *fiber.Ctx) error {
	var info structs.Fidelity
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	fidelity, err := f.m.Create(info)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(fidelity)

}

func (f FidelityHandler) HandleGet(ctx *fiber.Ctx) error {
	var info structs.Fidelity
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	fidelity, err := f.m.Read(info)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(fidelity)

}

func NewFidelidadeHandler(m middleware.FidelityMiddleware) FidelityHandler{
	return FidelityHandler{m: m}
}