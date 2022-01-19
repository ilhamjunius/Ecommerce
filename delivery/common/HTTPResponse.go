package common

import "net/http"

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//NewInternalServerErrorResponse default internal server error response
func NewSuccessOperationResponse() DefaultResponse {
	return DefaultResponse{
		http.StatusOK,
		"Successful Operation",
	}
}

//NewInternalServerErrorResponse default internal server error response
func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		http.StatusInternalServerError,
		"Internal Server Error",
	}
}

//NewNotFoundResponse default not found error response
func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		http.StatusNotFound,
		"Not Found",
	}
}

//NewBadRequestResponse default bad request error response
func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		http.StatusBadRequest,
		"Bad Request",
	}
}

//NewConflictResponse default conflict response error response
func NewConflictResponse() DefaultResponse {
	return DefaultResponse{
		http.StatusConflict,
		"Data Has Been Modified",
	}
}

//NewStatusNotAccepted default not
func NewStatusNotAcceptable() DefaultResponse {
	return DefaultResponse{
		http.StatusNotAcceptable,
		"Not Accepted",
	}
}

func NewStatusNotAuthorized() DefaultResponse {
	return DefaultResponse{
		http.StatusUnauthorized,
		"Not Authorized",
	}
}
