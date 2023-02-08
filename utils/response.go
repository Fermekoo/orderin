package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(code int, err error) gin.H {
	return gin.H{
		"code":    code,
		"message": err.Error(),
	}
}

type GeneralResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseOK(code int, message string, data interface{}) GeneralResponse {
	return GeneralResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
