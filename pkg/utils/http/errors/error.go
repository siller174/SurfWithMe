package errors

import "fmt"

type HTTPError interface {
	Error() string
	GetStatus() int
	ToResponse() string
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

func (err *Errors) ToResponse() string {
	return fmt.Sprintf("{ Error:\"%s\" }", err.Err.Error())
}

func NewRequestTimeout(err error) *Errors {
	return &Errors{
		Err:    err,
		Status: 408,
	}
}

func NewBadRequest(err error) *Errors {
	return &Errors{
		Err:    err,
		Status: 400,
	}
}

func NewInternalErr(err error) *Errors {
	return &Errors{
		Err:    err,
		Status: 500,
	}
}
