package handlers

import (
	"datav-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderItemUserCount(c *gin.Context) {
	var result struct {
		OrderCount int64 `json:"order_count"`
		ItemCount  int64 `json:"item_count"`
		UserCount  int64 `json:"user_count"`
	}
	query := `
		SELECT 
			(SELECT COUNT(DISTINCT order_id) FROM user_order_data) AS order_count,
			(SELECT COUNT(DISTINCT item_id) FROM user_order_data) AS item_count,
			(SELECT COUNT(DISTINCT user_id) FROM user_behavior) AS user_count
	`
	if err := db.DB.Get(&result, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
