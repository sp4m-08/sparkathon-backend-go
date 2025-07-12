package controllers

import (
	"encoding/json"
	"net/http"
	"upc-backend-sparkathon/main/models"
	"upc-backend-sparkathon/main/services"

	"github.com/gin-gonic/gin"
)

// set context for llm: barcode scanned -> using barcode no. fetch product details from db -> set these json details as context for llm
func ScanProductAndInitChat(c *gin.Context) {
	var req struct {
		UPC string `json:"upc"`
	}
	if err := c.BindJSON(&req); err != nil || req.UPC == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC"})
		return
	}

	var product models.Product
	if err := services.DB.Where("ean = ?", req.UPC).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Convert to map
	b, _ := json.Marshal(product)
	var productMap map[string]interface{}
	json.Unmarshal(b, &productMap)

	// Set context
	ctxResp, err := services.SetContext(map[string]interface{}{"product": productMap}, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LLM context failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": ctxResp.SessionID,
		"product":    product,
	})
}

func ChatAboutProduct(c *gin.Context) {
	var req struct {
		SessionID string `json:"session_id"`
		Message   string `json:"message"`
	}
	if err := c.BindJSON(&req); err != nil || req.Message == "" || req.SessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing message or session ID"})
		return
	}
	resp, err := services.Chat(req.Message, req.SessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LLM chat failed"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func GetChatHistory(c *gin.Context) {
	sessionID := c.Param("session_id")
	history, err := services.GetHistory(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch history"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"session_id": sessionID, "history": history})
}

func LLMHealthCheck(c *gin.Context) {
	status, err := services.CheckHealth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "LLM unavailable", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, status)
}

// func ScanAndSaveProduct(c *gin.Context) {
// 	var request struct {
// 		UPC string `json:"upc"`
// 	}

// 	if err := c.BindJSON(&request); err != nil || request.UPC == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing UPC"})
// 		return
// 	}

// 	product, err := services.FetchProductfromUPC(request.UPC)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := services.DB.Create(product).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "product saved successfully",
// 		"product": product,
// 	})
// }

// func GetProducts(c *gin.Context) {
// 	var ProductList []models.Product
// 	if err := services.DB.Find(&ProductList).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, ProductList)
// }

// func GetProduct(c *gin.Context) {
// 	ean := c.Param("ean")

// 	var thisproduct models.Product
// 	if err := services.DB.Where("ean = ?", ean).First(&thisproduct).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, thisproduct)
// }
