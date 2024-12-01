package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdsProductAnalysis(c *gin.Context) {
	var summary models.AdsProductAnalysisSummary
	query := `
		SELECT
			SUM(browse_cnt) AS browse_cnt,
			SUM(like_cnt) AS like_cnt,
			SUM(add_cart_cnt) AS add_cart_cnt,
			SUM(purchase_cnt) AS purchase_cnt
		FROM
			ads_product_analysis
	`
	if err := db.DB.Get(&summary, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
