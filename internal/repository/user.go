package repository

import (
	"context"
	"github.com/nicitapa/firstProgect/internal/models"
	"github.com/rs/zerolog"
	"os"
)

func (r *Repository) CreateUser(ctx context.Context, user models.User) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateUser").Logger()
	_, err = r.db.ExecContext(ctx, `INSERT INTO users (full_name, username, password, role)
					VALUES ($1, $2, $3, $4)`,
		user.FullName,
		user.Username,
		user.Password,
		user.Role)
	if err != nil {
		logger.Err(err).Msg("error inserting user")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (user models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetUserByID").Logger()
	if err = r.db.GetContext(ctx, &user, `
		SELECT id, full_name, username, password, role, created_at, updated_at 
		FROM users
		WHERE id = $1`, id); err != nil {
		r.logger.Error().Err(err).Str("func", "repository.GetUserByID").Msg("Error selecting users")
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (user models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetUserByUsername").Logger()
	if err = r.db.GetContext(ctx, &user, `
		SELECT id, full_name, username, password, role, created_at, updated_at 
		FROM users
		WHERE username = $1`, username); err != nil {
		logger.Err(err).Msg("error selecting user")
		return models.User{}, r.translateError(err)
	}

	return user, nil
}
