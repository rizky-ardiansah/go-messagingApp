package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-messagingApp/app/models"
	"github.com/rizky-ardiansah/go-messagingApp/app/repository"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	err := ctx.BodyParser(user)
	if err != nil {
		errorResponse := fmt.Errorf("failed to parse request: %v", err)
		fmt.Println(errorResponse)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = user.Validate()
	if err != nil {
		errorResponse := fmt.Errorf("failed to validate request: %v", err)
		fmt.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errorResponse.Error(), nil)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		errorResponse := fmt.Errorf("failed to encrypt password: %v", err)
		fmt.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errorResponse.Error(), nil)
	}

	user.Password = string(hashPassword)

	err = repository.InsertNewUser(ctx.Context(), user)
	if err != nil {
		errorResponse := fmt.Errorf("failed to create user: %v", err)
		fmt.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errorResponse.Error(), nil)
	}

	resp := user
	resp.Password = ""

	return response.SendSuccessResponse(ctx, resp)
}
