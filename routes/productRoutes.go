package routes

import (
	"upc-backend-sparkathon/main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	api := r.Group("/products")
	{
		//api.POST("/scan", controllers.ScanAndSaveProduct)
		// api.GET("/fetch", controllers.GetProducts)
		// api.GET("/fetch/:ean", controllers.GetProduct)

		api.POST("/scan", controllers.ScanProductAndInitChat)
		api.POST("/chat", controllers.ChatAboutProduct)
		api.GET("/chat/history/:session_id", controllers.GetChatHistory)
		api.GET("/chat/health", controllers.LLMHealthCheck)
	}

}
