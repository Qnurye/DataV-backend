package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProvinceDistribution(c *gin.Context) {
	var distributions []models.ProvinceDistribution
	query := "SELECT * FROM province_distribution"
	if err := db.DB.Select(&distributions, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, distributions)
}
