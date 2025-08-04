package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rizky-ardiansah/go-messagingApp/app/models"
	"github.com/rizky-ardiansah/go-messagingApp/app/repository"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/jwt_token"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/response"
	"go.elastic.co/apm"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	span, spanCtx := apm.StartSpan(ctx.Context(), "Register", "controller")
	defer span.End()

	user := new(models.User)

	err := ctx.BodyParser(user)
	if err != nil {
		errorResponse := fmt.Errorf("failed to parse request: %v", err)
		log.Println(errorResponse)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = user.Validate()
	if err != nil {
		errorResponse := fmt.Errorf("failed to validate request: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errorResponse.Error(), nil)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		errorResponse := fmt.Errorf("failed to encrypt password: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errorResponse.Error(), nil)
	}

	user.Password = string(hashPassword)

	err = repository.InsertNewUser(spanCtx, user)
	if err != nil {
		errorResponse := fmt.Errorf("failed to create user: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, errorResponse.Error(), nil)
	}

	resp := user
	resp.Password = ""

	return response.SendSuccessResponse(ctx, resp)
}

func Login(ctx *fiber.Ctx) error {
	span, spanCtx := apm.StartSpan(ctx.Context(), "Login", "controller")
	defer span.End()
	// parsing req dan validasi req
	var (
		loginReq = new(models.LoginRequest)
		resp     models.LoginResponse
		now      = time.Now()
	)

	err := ctx.BodyParser(loginReq)
	if err != nil {
		errorResponse := fmt.Errorf("failed to parse request: %v", err)
		log.Println(errorResponse)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = loginReq.Validate()
	if err != nil {
		errorResponse := fmt.Errorf("failed to validate request: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusBadRequest, errorResponse.Error(), nil)
	}

	user, err := repository.GetUserByUsername(spanCtx, loginReq.Username)
	if err != nil {
		errorResponse := fmt.Errorf("failed to get username: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusNotFound, "username/password salah", nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		errorResponse := fmt.Errorf("failed to check password: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusNotFound, "username/password salah", nil)
	}

	token, err := jwt_token.GenerateToken(spanCtx, user.Username, user.FullName, "token", now)
	if err != nil {
		errorResponse := fmt.Errorf("failed to generate token: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "terjadi kesalahan pada sistem", nil)
	}
	refreshToken, err := jwt_token.GenerateToken(spanCtx, user.Username, user.FullName, "refresh_token", now)
	if err != nil {
		errorResponse := fmt.Errorf("failed to generate token: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "terjadi kesalahan pada sistem", nil)
	}

	userSession := &models.UserSession{
		UserID:              user.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(jwt_token.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(jwt_token.MapTypeToken["refresh_token"]),
	}
	err = repository.InsertNewUserSession(spanCtx, userSession)
	if err != nil {
		errorResponse := fmt.Errorf("failed insert user session: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "terjadi kesalahan pada sistem", nil)
	}

	resp.Username = user.Username
	resp.FullName = user.FullName
	resp.Token = token
	resp.RefreshToken = refreshToken

	return response.SendSuccessResponse(ctx, resp)
}

func Logout(ctx *fiber.Ctx) error {
	span, spanCtx := apm.StartSpan(ctx.Context(), "Logout", "controller")
	defer span.End()
	token := ctx.Get("Authorization")
	err := repository.DeleteUserSessionByToken(spanCtx, token)
	if err != nil {
		errorResponse := fmt.Errorf("failed to delete user session: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "terjadi kesalahan pada sistem", nil)
	}
	return response.SendSuccessResponse(ctx, nil)
}

func RefreshToken(ctx *fiber.Ctx) error {
	span, spanCtx := apm.StartSpan(ctx.Context(), "RefreshToken", "controller")
	defer span.End()
	now := time.Now()
	refreshToken := ctx.Get("Authorization")
	username := ctx.Locals("username").(string)
	fullName := ctx.Locals("full_name").(string)

	token, err := jwt_token.GenerateToken(spanCtx, username, fullName, "token", now)
	if err != nil {
		errorResponse := fmt.Errorf("failed to generate token: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "terjadi kesalahan pada sistem", nil)
	}

	err = repository.UpdateUserSessionToken(spanCtx, token, now.Add(jwt_token.MapTypeToken["token"]), refreshToken)
	if err != nil {
		errorResponse := fmt.Errorf("failed to update token: %v", err)
		log.Println(errorResponse)
		return response.SendFailureResponse(ctx, fiber.StatusInternalServerError, "terjadi kesalahan pada sistem", nil)
	}
	return response.SendSuccessResponse(ctx, fiber.Map{
		"token": token,
	})
}
