package utils

import "errors"

var (
	ErrNotFound = errors.New("ES not found")
	ErrConflict = errors.New("ES conflict")
)
