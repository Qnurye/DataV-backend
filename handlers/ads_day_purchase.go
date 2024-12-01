package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdsDayPurchase(c *gin.Context) {
	var purchases []models.AdsDayPurchase
	query := "SELECT dt, purchase_cnt FROM ads_day_purchase order by dt"
	if err := db.DB.Select(&purchases, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, purchases)
}
