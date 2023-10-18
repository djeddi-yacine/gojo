package db

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

var ErrRecordNotFound = pgx.ErrNoRows

var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}

var ErrForeignKeyViolation = &pgconn.PgError{
	Code: ForeignKeyViolation,
}

func DatabaseError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case UniqueViolation:
			return ErrUniqueViolation
		case ForeignKeyViolation:
			return ErrForeignKeyViolation
		default:
			return fmt.Errorf("database Error: %s - %s", pgErr.Code, pgErr.Message)
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrRecordNotFound
	}

	return fmt.Errorf("database Error: %s", err.Error())
}

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}

func ErrorSQL(err error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		log.Err(pgErr).Str("SQL Message:", pgErr.Message)
		log.Err(pgErr).Str("SQL Code:", pgErr.Code)
	}
}
