package routes

import (
	"upc-backend-sparkathon/main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	api := r.Group("/products")
	{
		//api.POST("/scan", controllers.ScanAndSaveProduct)
		api.GET("/fetch", controllers.GetProducts)
		api.GET("/fetch/:ean", controllers.GetProduct)
	}

}
