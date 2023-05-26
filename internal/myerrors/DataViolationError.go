package myerrors

import (
	"fmt"
)

type DataViolationError struct {
	err  error
	name string
}

func (dv *DataViolationError) Error() string {
	return fmt.Sprintf("data with name %s already exists", dv.name)
}

func (dv *DataViolationError) Unwrap() error {
	return dv.err
}

func NewDataViolationError(name string, err error) error {
	return &DataViolationError{
		name: name,
		err:  err,
	}
}
