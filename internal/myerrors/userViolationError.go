package myerrors

import (
	"fmt"
)

type UserViolationError struct {
	err   error
	login string
}

func (ve *UserViolationError) Error() string {
	return fmt.Sprintf("user with login %s already exists", ve.login)
}

func (ve *UserViolationError) Unwrap() error {
	return ve.err
}

func NewUserViolationError(login string, err error) error {
	return &UserViolationError{
		login: login,
		err:   err,
	}
}
