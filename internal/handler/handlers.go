package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jumuia/akademy/internal/service"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	FetchUser(c *gin.Context)
	FetchAllUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewUserHandler(svc service.UserService) UserHandler {
	return &userHandler{
		Svc: svc,
	}
}

type errResponse struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

type userHandler struct {
	Svc service.UserService
}

// FetchAllUsers implements UserHandler
func (*userHandler) FetchAllUsers(c *gin.Context) {
	panic("unimplemented")
}

// FetchUser implements UserHandler
func (h *userHandler) FetchUser(c *gin.Context) {
	u := fetchUserRequest{}
	err := c.ShouldBindUri(&u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse{
			ErrorCode:    "ERR1000002",
			ErrorMessage: "ID is mandatory",
		})
		return
	}
	user, err := h.Svc.FindUser(u.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse{
			ErrorCode:    "ERR1000003",
			ErrorMessage: "Unable to fetch user",
		})
		return
	}
	userResp := createUserRequest{
		FirstName:   user.FirstName,
		MiddleName:  user.MiddleName,
		LastName:    user.LastName,
		DOB:         user.DOB,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
	c.JSONP(http.StatusOK, userResp)
}

// CreateUser implements UserHandler
func (h *userHandler) CreateUser(c *gin.Context) {
	u := createUserRequest{}
	c.Bind(&u)

	user := service.User{
		FirstName:   u.FirstName,
		MiddleName:  u.MiddleName,
		LastName:    u.LastName,
		DOB:         u.DOB,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}
	err := h.Svc.Create(&user)
	if err != nil {
		// handle error
		log.Printf("unable to create user: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse{
			ErrorCode:    "ERR1000001",
			ErrorMessage: "Unable to create the user",
		})
		return
	}
	c.JSONP(http.StatusCreated, gin.H{
		"status": "created user",
	})
	return
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
