package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserBehavior(c *gin.Context) {
	var behaviors []models.UserBehavior
	query := "SELECT * FROM user_behavior"
	if err := db.DB.Select(&behaviors, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, behaviors)
}
