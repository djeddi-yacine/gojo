package db

import (
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

func ErrorSQL(err error) {
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
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
			Message: "Not Found",
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

func ErrorDBType(err error) string {
	if err != nil {
		if err == pgx.ErrNoRows {
			return "Not found"
		}
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return "Already exist."
			case pgerrcode.ForeignKeyViolation:
				return "Others data have relation with this model."
			}
			if pgerrcode.IsCaseNotFound(pgErr.Code) {
				return "Not found."
			}
			if pgerrcode.IsConnectionException(pgErr.Code) {
				return "Database connection timeout."
			}
			if pgerrcode.IsDataException(pgErr.Code) {
				return "Data exception."
			}
			if pgerrcode.IsInvalidTransactionState(pgErr.Code) {
				return "Failed to add it to database."
			}
		}
	}

	return ""
}
