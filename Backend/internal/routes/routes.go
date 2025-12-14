package routes

import (
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/handlers"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Public Sweets routes
	sweetsPublic := api.Group("/sweets")
	{
		sweetsPublic.GET("", handlers.ListSweets)
		sweetsPublic.GET("/search", handlers.SearchSweets)
	}

	// Protected Sweets routes
	sweets := api.Group("/sweets")
	sweets.Use(middleware.AuthMiddleware())
	{
		sweets.POST("", middleware.AdminMiddleware(), handlers.AddSweet)
		sweets.PUT("/:id", middleware.AdminMiddleware(), handlers.UpdateSweet)
		sweets.DELETE("/:id", middleware.AdminMiddleware(), handlers.DeleteSweet)

		// Inventory routes
		sweets.POST("/:id/purchase", handlers.PurchaseSweet)
		sweets.POST("/:id/restock", middleware.AdminMiddleware(), handlers.RestockSweet)
	}

	// Admin routes
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/transactions", handlers.GetAllTransactions)
	}
}
