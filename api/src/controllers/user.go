package controllers

import (
	"encoding/json"
	"net/http"
	grpc_user "sword-health/api/grpc/user"
	"sword-health/api/infra/amqp"

	"sword-health/api/validators"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Validator  *validators.JSONValidator
	UserClient *grpc_user.UserClient
	AMQ        *amqp.Connection
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

	if message, err := json.Marshal(request); err == nil {
		u.AMQ.Dispatch(amqp.ExchangeUser, amqp.RouteKeyUserCreate, message)
	}

	c.JSON(http.StatusAccepted, http.StatusText(http.StatusAccepted))
}

func (u *UserController) Update(c *gin.Context) {

	var request CreateRequest

	if errors := u.Validator.Validate(c, &request); errors != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "eerr"})
}

func (u *UserController) Delete(c *gin.Context) {
	// id := c.Param("id")
}
