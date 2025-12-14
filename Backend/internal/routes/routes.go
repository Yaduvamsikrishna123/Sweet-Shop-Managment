package routes

import (
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/handlers"
	"github.com/Yaduvamsikrishna123/Sweet-Shop-Management-System/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handlers.Handler) {
	api := r.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	// Public Sweets routes
	sweetsPublic := api.Group("/sweets")
	{
		sweetsPublic.GET("", h.ListSweets)
		sweetsPublic.GET("/search", h.SearchSweets)
	}

	// Protected Sweets routes
	sweets := api.Group("/sweets")
	sweets.Use(middleware.AuthMiddleware())
	{
		sweets.POST("", middleware.AdminMiddleware(), h.AddSweet)
		sweets.PUT("/:id", middleware.AdminMiddleware(), h.UpdateSweet)
		sweets.DELETE("/:id", middleware.AdminMiddleware(), h.DeleteSweet)

		// Inventory routes
		sweets.POST("/:id/purchase", h.PurchaseSweet)
		sweets.POST("/:id/restock", middleware.AdminMiddleware(), h.RestockSweet)
	}

	// Admin routes
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/transactions", h.GetAllTransactions)
	}
}
