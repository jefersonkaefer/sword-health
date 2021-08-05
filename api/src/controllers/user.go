package controllers

import (
	"net/http"
	grpc_user "sword-health/api/grpc/user"

	"sword-health/api/validators"

	"github.com/gin-gonic/gin"
	status "google.golang.org/grpc/status"
)

type UserController struct {
	Validator  *validators.JSONValidator
	UserClient *grpc_user.UserClient
}

type CreateRequest struct {
	FirstName  string `json:"first_name" binding:"required" validate:"min=2,max=32,alpha"`
	LastName   string `json:"last_name" binding:"required" validate:"min=2,max=32,alpha"`
	Email      string `json:"email" binding:"required" validate:"required,email"`
	Password   string `json:"password" binding:"required" validate:"min=8,max=32,alphanum"`
	Repassword string `json:"confirm_password" binding:"required" validate:"eqfield=Password,required"`
	Role       string `json:"role" binding:"required" validate:"required"`
}

func (u *UserController) Create(c *gin.Context) {

	var request CreateRequest

	if errors := u.Validator.Validate(c, &request); errors != nil {
		return
	}

	user, err := u.UserClient.CreateUser(
		request.Email,
		request.Password,
		request.Repassword,
		request.FirstName,
		request.LastName,
		request.Role,
	)

	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			c.JSON(http.StatusBadGateway, gin.H{"error": http.StatusText(http.StatusBadGateway)})
			return
		}
		c.JSON(int(status.Code()), gin.H{"error": status.Message()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": user.GetId()})
}
