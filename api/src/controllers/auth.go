package controllers

import (
	"net/http"
	"os"
	grpc_user "sword-health/api/grpc/user"
	"time"

	"sword-health/api/validators"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	status "google.golang.org/grpc/status"
)

type AuthController struct {
	Validator  *validators.JSONValidator
	UserClient *grpc_user.UserClient
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

func (a *AuthController) Login(c *gin.Context) {

	var request LoginRequest
	var token string

	if errors := a.Validator.Validate(c, &request); errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	user, err := a.UserClient.Login(request.Email, request.Password)

	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			c.JSON(http.StatusBadGateway, gin.H{"error": http.StatusText(http.StatusBadGateway)})
			return
		}
		c.JSON(int(status.Code()), gin.H{"error": status.Message()})
		return
	}

	token, err = a.CreateToken(
		user.GetId(),
		user.GetEmail(),
		user.GetFirstName(),
		user.GetLastName(),
		user.GetRole(),
	)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) CreateToken(
	userid int32,
	email string,
	firstName string,
	lastName string,
	role string,
) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}

	atClaims["user_id"] = userid
	atClaims["email"] = email
	atClaims["first_name"] = firstName
	atClaims["last_name"] = userid
	atClaims["role"] = role

	atClaims["authorized"] = true
	atClaims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
