package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTop5PopularItems(c *gin.Context) {
	var items []models.AdsProductAnalysis
	query := `
		SELECT item_id, browse_cnt, like_cnt, purchase_cnt, popularity_score
		FROM ads_product_analysis
		ORDER BY popularity_score DESC
		LIMIT 5
	`
	if err := db.DB.Select(&items, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
