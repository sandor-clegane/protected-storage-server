package myerrors

import (
	"fmt"
)

type InvalidPasswordError struct {
	password string
}

func (ip *InvalidPasswordError) Error() string {
	return fmt.Sprintf("password %s invalid", ip.password)
}

func NewInvalidPasswordError(password string) error {
	return &InvalidPasswordError{
		password: password,
	}
}
