package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kooroshh/fiber-boostrap/app/models"
	"github.com/kooroshh/fiber-boostrap/app/repository"
)

func Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	err := ctx.BodyParser(user)
	if err != nil {
		fmt.Println("Failed to parse request: ", err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = user.Validate()
	if err != nil {
		fmt.Println("Failed to validate request: ", err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = repository.InsertNewUser(ctx.Context(), user)
	if err != nil {
		fmt.Println("Failed to create user", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
