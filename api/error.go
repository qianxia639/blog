package api

import "database/sql"

var (
	ErrNoRows             = sql.ErrNoRows
	ErrUniqueViolation    = "unique_violation"
	ErrForeignKyViolation = "foreign_key_violation"
	ErrInvalidParameter   = "invalid paramparameter"
)
