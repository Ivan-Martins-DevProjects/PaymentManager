package apperrors

type AppError interface {
	error
	Code() string
	Message() string
}

type InternalError struct {
	InternalCode    string
	InternalMessage string
	TracebackError  error
}

func (e *InternalError) Error() string {
	if e.TracebackError != nil {
		return e.InternalCode + ": " + e.TracebackError.Error()
	}
	return e.InternalCode
}

func (e *InternalError) Unrwap() error {
	return e.TracebackError
}

func (e *InternalError) Code() string {
	return e.InternalCode
}

func (e *InternalError) Message() string {
	return e.InternalMessage
}

func (e *InternalError) Is(target error) bool {
	t, ok := target.(*InternalError)
	return ok && e.InternalCode == t.InternalCode
}

func GenerateError(errorVar InternalError, err error) *InternalError {
	errorVar.TracebackError = err
	return &errorVar
}
