package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-messagingApp/app/repository"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/response"
)

func GetHistory(ctx *fiber.Ctx) error {
	resp, err := repository.GetAllMessage(ctx.Context())
	if err != nil {
		fmt.Println(err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "Terjadi kesalahan pada server", nil)
	}
	return response.SendSuccessResponse(ctx, resp)
}