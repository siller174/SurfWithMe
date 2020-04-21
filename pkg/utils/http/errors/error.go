package errors

import (
	"fmt"
	"net/http"
)

type HTTPError interface {
	Error() string
	GetStatus() int
	ToResponse(bool) string
}

type Errors struct {
	Err    error
	Status int
}

func (err *Errors) Error() string {
	return err.Err.Error()
}

func (err *Errors) GetStatus() int {
	return err.Status
}

func (err *Errors) ToResponse(addErr bool) string {
	if addErr {
		return fmt.Sprintf("{\nError:\"%s\"\n}", err.Err.Error())
	} else {
		return fmt.Sprintf("{\nError:\"Something was wrong\"\n}")

	}
}

func NewBadRequest(err error) *Errors {
	return &Errors{
		Err:    err,
		Status: http.StatusBadRequest,
	}
}

func NewInternalErr(err error) *Errors {
	return &Errors{
		Err:    err,
		Status: http.StatusInternalServerError,
	}
}

func NewNotFound(err error) *Errors {
	return &Errors{
		Err:    err,
		Status: http.StatusNotFound,
	}
}
