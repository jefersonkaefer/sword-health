package controllers

import (
	"net/http"
	"strconv"
	grpc_notification "sword-health/api/grpc/notification"
	middleware "sword-health/api/http"

	"sword-health/api/validators"

	"github.com/gin-gonic/gin"
	status "google.golang.org/grpc/status"
)

type NotificationController struct {
	Validator          *validators.JSONValidator
	NotificationClient *grpc_notification.NotificationClient
}

func (u *NotificationController) Get(c *gin.Context) {

	userLogged, _ := c.Get("userLogged")
	user := userLogged.(*middleware.UserLogged)

	notificationId := c.Param("id")
	id, _ := strconv.Atoi(notificationId)

	notification, err := u.NotificationClient.Get(int(user.Id), id)

	if err != nil {
		status, ok := status.FromError(err)
		if !ok {
			c.JSON(http.StatusBadGateway, gin.H{"error": http.StatusText(http.StatusBadGateway)})
			return
		}
		c.JSON(int(status.Code()), gin.H{"error": status.Message()})
		return
	}
	c.JSON(http.StatusOK, notification)
}

func (u *NotificationController) List(c *gin.Context) {
	userLogged, _ := c.Get("userLogged")
	user := userLogged.(*middleware.UserLogged)

	from := c.Query("from_id")
	fromId, _ := strconv.Atoi(from)

	limitParam := c.Query("limit")
	limit, _ := strconv.Atoi(limitParam)

	notifications, err := u.NotificationClient.List(
		int(user.Id),
		fromId,
		limit,
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

	c.JSON(http.StatusOK, notifications)

}
