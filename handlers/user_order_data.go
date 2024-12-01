package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserOrderData(c *gin.Context) {
	var orders []models.UserOrderData
	query := "SELECT * FROM user_order_data"
	if err := db.DB.Select(&orders, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
