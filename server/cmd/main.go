package main

import (
	"store-app/server/internal/config"
	"store-app/server/internal/database"
	"store-app/server/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	database.Init(cfg)
	database.Seed(cfg.BusinessType)

	r := gin.Default()

	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/store", handler.GetStore)
		api.PUT("/store", handler.UpdateStore)

		api.GET("/categories", handler.GetCategories)
		api.POST("/categories", handler.CreateCategory)
		api.PUT("/categories/:id", handler.UpdateCategory)
		api.DELETE("/categories/:id", handler.DeleteCategory)

		api.GET("/products", handler.GetProducts)
		api.GET("/products/:id", handler.GetProduct)
		api.POST("/products", handler.CreateProduct)
		api.PUT("/products/:id", handler.UpdateProduct)
		api.DELETE("/products/:id", handler.DeleteProduct)

		api.GET("/cart", handler.GetCart)
		api.POST("/cart", handler.AddToCart)
		api.PUT("/cart/:id", handler.UpdateCartItem)
		api.DELETE("/cart/:id", handler.DeleteCartItem)

		api.POST("/orders", handler.CreateOrder)
		api.GET("/orders", handler.GetOrders)
		api.GET("/orders/:id", handler.GetOrder)
	}

	r.Run(":" + cfg.ServerPort)
}
