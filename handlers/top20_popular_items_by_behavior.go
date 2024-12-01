package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTop20PopularItemsByBehavior(c *gin.Context) {
	var items []models.Top20PopularItemsByBehavior
	query := "SELECT * FROM top_20_popular_items_by_behavior order by popularity desc"
	if err := db.DB.Select(&items, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
