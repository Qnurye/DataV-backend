package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProvinceOrders(c *gin.Context) {
	var orders []models.ProvinceOrder
	query := "SELECT province_id, province, orders_cnt FROM ads_province_orders_display"

	if err := db.DB.Select(&orders, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
