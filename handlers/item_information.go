package handlers

import (
	"datav-api/db"
	"datav-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItemInformation(c *gin.Context) {
	var items []models.ItemInformation
	query := "SELECT * FROM item_information"
	if err := db.DB.Select(&items, query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
