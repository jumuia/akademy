package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jumuia/akademy/internal/service"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	ReadUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserHandler(svc service.UserService) UserHandler {
	return &userHandler{
		Svc: svc,
	}
}

type userHandler struct {
	Svc service.UserService
}

// CreateUser implements UserHandler
func (*userHandler) CreateUser(c *gin.Context) {
	panic("unimplemented")
}

// DeleteUser implements UserHandler
func (*userHandler) DeleteUser(c *gin.Context) {
	panic("unimplemented")
}

// ReadUser implements UserHandler
func (*userHandler) ReadUser(c *gin.Context) {
	panic("unimplemented")
}

// UpdateUser implements UserHandler
func (*userHandler) UpdateUser(c *gin.Context) {
	panic("unimplemented")
}
