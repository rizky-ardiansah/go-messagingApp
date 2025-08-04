package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-messagingApp/app/repository"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/response"
	"go.elastic.co/apm"
)

func GetHistory(ctx *fiber.Ctx) error {
	span, spanCtx := apm.StartSpan(ctx.Context(), "GetHistory", "controller")
	defer span.End()
	resp, err := repository.GetAllMessage(spanCtx)
	if err != nil {
		log.Println(err)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "Terjadi kesalahan pada server", nil)
	}
	return response.SendSuccessResponse(ctx, resp)
}
