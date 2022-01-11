package models

var (
	NotFoundError        = NewAppError(nil, 404, "Not Found")
	ConnectionError      = NewAppError(nil, 503, "Connection Error")
	ExistFoundError      = NewAppError(nil, 409, "Item is already exists")
	PasswordOrLoginError = NewAppError(nil, 403, "Password or Login is incorrect")
)

type AppError struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	Message    string `json:"messsage"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(err error, StatusCode int, Message string) *AppError {
	return &AppError{
		Err:        err,
		Message:    Message,
		StatusCode: StatusCode,
	}
}
