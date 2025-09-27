package service

import (
	"context"
	"errors"
	"github.com/nicitapa/firstProgect/internal/errs"
	"github.com/nicitapa/firstProgect/internal/models"
	"github.com/nicitapa/firstProgect/utils"
)

func (s *Service) CreateUser(ctx context.Context, user models.User) (err error) {
	// Проверить существует ли пользователь с таким username'ом в бд
	_, err = s.repository.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return err
		}
	} else {
		return errs.ErrUsernameAlreadyExists
	}

	// За хэшировать пароль
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return err
	}

	user.Role = models.RoleUser

	// Добавить пользователя в бд
	if err = s.repository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *Service) Authenticate(ctx context.Context, user models.User) (int, models.Role, error) {
	// проверить существует ли пользователь с таким username
	userFromDB, err := s.repository.GetUserByUsername(ctx, user.Username)
	if err != nil {
		if !errors.Is(err, errs.ErrNotfound) {
			return 0, "", errs.ErrUserNotFound
		}

		return 0, "", err
	}

	// за хэшировать пароль, который получили от пользователя
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return 0, "", err
	}

	// проверить правильно ли он указал пароль
	if userFromDB.Password != user.Password {
		return 0, "", errs.ErrIncorrectUsernameOrPassword
	}

	return userFromDB.ID, userFromDB.Role, nil
}
