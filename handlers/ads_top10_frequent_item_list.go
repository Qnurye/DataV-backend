package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTop10FrequentItemList(c *gin.Context) {
	var items []models.AdsTop10FrequentItemList
	query := "SELECT item_list, cnt FROM ads_top10_frequent_item_list"
	if err := db.DB.Select(&items, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
