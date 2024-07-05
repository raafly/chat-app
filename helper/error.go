package helper

type ErrResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Content interface{} `json:"data"`
}

func (e *ErrResponse) Error() string {
	return e.Message
}

func (e *Response) Error() string {
	return e.Message
}

func NewCreated(message string) error {
	return &ErrResponse{
		Code:    201,
		Status:  true,
		Message: message,
	}
}

func NewContent(data any) error {
	return &Response{
		Code: 200,
		Status: true,
		Message: "SUCCESS",
		Content: data,
	}
}

func NewSucces(message string) error {
	return &ErrResponse{
		Code:    200,
		Status:  true,
		Message: message,
	}
}

func ErrBadRequest(message string) error {
	return &ErrResponse{
		Code:    404,
		Status:  false,
		Message: message,
	}
}

func ErrNotFound(message string) error {
	return &ErrResponse{
		Code:    404,
		Status:  false,
		Message: message,
	}
}

func ErrInternalServerError() error {
	return &ErrResponse{
		Code:    500,
		Status:  false,
		Message: "INTERNAL SERVER ERROR",
	}
}
