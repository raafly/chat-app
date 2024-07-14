package helper

type ErrResponse struct {
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (e *ErrResponse) Error() string {
	return e.Message
}

func NewCreated(message string, data any) error {
	return &ErrResponse{
		Status:  true,
		Message: message,
		Data: data,
	}
}

func NewContent(message string, data any) error {
	return &ErrResponse{
		Status:  true,
		Message: message,
		Data: data,
	}
}

func NewSucces(message string, data any) error {
	return &ErrResponse{
		Status:  true,
		Message: message,
		Data: data,
	}
}

func ErrBadRequest(message string, data any) error {
	return &ErrResponse{
		Status:  false,
		Message: message,
		Data: data,
	}
}

func ErrNotFound(message string, data any) error {
	return &ErrResponse{
		Status:  false,
		Message: message,
		Data: data,
	}
}

func ErrInternalServerError() error {
	return &ErrResponse{
		Status:  false,
		Message: "INTERNAL SERVER ERROR",
		Data: nil,
	}
}
