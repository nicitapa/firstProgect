package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/nicitapa/firstProgect/internal/errs"
	"github.com/rs/zerolog"
)

type Repository struct {
	db     *sqlx.DB
	logger zerolog.Logger
}

func NewRepository(db *sqlx.DB, logger zerolog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger}

}
func (r *Repository) translateError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return errs.ErrNotfound
	default:
		return err
	}
}
