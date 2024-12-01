package main

import (
	"datav-api/config"
	"datav-api/db"
	"datav-api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}))

	r.GET("/province", handlers.GetProvinceOrders)
	r.GET("/ads_day_purchase", handlers.GetAdsDayPurchase)
	r.GET("/ads_product_analysis", handlers.GetAdsProductAnalysis)
	r.GET("/ads_top10_frequent_item_list", handlers.GetTop10FrequentItemList)
	r.GET("/item_information", handlers.GetItemInformation)
	r.GET("/province_distribution", handlers.GetProvinceDistribution)
	r.GET("/top20_popular_items_by_behavior", handlers.GetTop20PopularItemsByBehavior)
	r.GET("/user_behavior", handlers.GetUserBehavior)
	r.GET("/user_order_data", handlers.GetUserOrderData)
	r.GET("/ads_top5_popular_items", handlers.GetTop5PopularItems)
	r.GET("/order_item_user_count", handlers.GetOrderItemUserCount)
	err := r.Run(":2026")
	if err != nil {
		return
	}
}
