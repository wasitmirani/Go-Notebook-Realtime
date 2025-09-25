package services

import (
	"GoNotebookRealtime/internal/app/models"
	"GoNotebookRealtime/internal/app/repositories"
	"GoNotebookRealtime/internal/pkg/utils"
	"errors"
)

func Login(email, password string) (*models.User, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
