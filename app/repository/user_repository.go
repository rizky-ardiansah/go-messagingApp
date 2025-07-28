package repository

import (
	"context"

	"github.com/rizky-ardiansah/go-messagingApp/app/models"
	"github.com/rizky-ardiansah/go-messagingApp/pkg/database"
)

func InsertNewUser(ctx context.Context, user *models.User) error {
	return database.DB.Create(user).Error
}

func InsertNewUserSession(ctx context.Context, session *models.UserSession) error {
	return database.DB.Create(session).Error
}

func GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var (
		resp models.User
		err  error
	)
	err = database.DB.Where("username = ?", username).Last(&resp).Error
	return resp, err
}
