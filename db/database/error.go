package db

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

var (
	ErrInccorectPassword  = errors.New("inccorect password")
	ErrFailedPrecondition = errors.New("this device is banned for this account")
)

func ErrorSQL(err error) {
	if err != nil {
		pgErr := ErrorDB(err)
		if pgErr != nil {
			log.Error().Msgf("SQL Message: %v", pgErr.Message)
			log.Error().Msgf("SQL Code: %v", pgErr.Code)
		}
	}
}

func ErrorDB(err error) *pgconn.PgError {
	var pgErr *pgconn.PgError
	if err == pgx.ErrNoRows {
		return &pgconn.PgError{
			Code:    pgerrcode.CaseNotFound,
			Message: "not found",
			Detail:  err.Error(),
		}
	}
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			return pgErr
		}
	}

	return pgErr
}

func ErrorType(err error) string {
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return "not found"
		case ErrInccorectPassword:
			return err.Error()
		case ErrFailedPrecondition:
			return err.Error()
		}

		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return "already exist."
			case pgerrcode.ForeignKeyViolation:
				return "others data have relation with this model."
			}
			if pgerrcode.IsCaseNotFound(pgErr.Code) {
				return "not found."
			}
			if pgerrcode.IsConnectionException(pgErr.Code) {
				return "database connection timeout."
			}
			if pgerrcode.IsDataException(pgErr.Code) {
				return "data exception."
			}
			if pgerrcode.IsInvalidTransactionState(pgErr.Code) {
				return "failed to add it to database."
			}
		}
	}

	return ""
}
