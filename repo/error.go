package repo

import "errors"

var (
	ErrNotFound = errors.New("Not_Found")
	ErrGeneral  = errors.New("DB_Error")
)
