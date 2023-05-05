package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jumuia/akademy/internal/service"
)

func NewUserHandler(svc service.UserService) UserHandler {
	return &userHandler{
		Svc: svc,
	}
}

type errResponse struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	FetchUser(c *gin.Context)
	FetchAllUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
