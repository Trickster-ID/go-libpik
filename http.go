package libpik

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObject struct{}

//this func for build success response data
func BSuccessResponse(data interface{}) Response {
	res := Response{
		Status:  true,
		Message: "OK",
		Error:   nil,
		Data:    data,
	}
	return res
}

//this func for build fail response data
func BErrorResponse(message string, err string) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   strings.Split(err, "\n"),
		Data:    EmptyObject{},
	}
	return res
}
