package loader

type DBLoadError struct {
	msg string
}

func (e *DBLoadError) Error() string {
	return e.msg
}

func NewDBLoadError(msg string) error {
	return &DBLoadError{
		msg: msg,
	}
}
