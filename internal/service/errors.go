package service

type InvalidCountryError struct {
	msg string
}

func (e *InvalidCountryError) Error() string {
	return e.msg
}

func NewInvalidCountryError(msg string) error {
	return &InvalidCountryError{
		msg: msg,
	}
}
