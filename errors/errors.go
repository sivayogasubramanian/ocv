package ocverrs

type Error interface {
	StatusCode() int
	Message() string
}

type OcvError struct {
	statusCode int
	message    string
}

func New(statusCode int, message string) *OcvError {
	return &OcvError{statusCode: statusCode, message: message}
}

func (e *OcvError) StatusCode() int {
	return e.statusCode
}

func (e *OcvError) Message() string {
	return e.message
}
