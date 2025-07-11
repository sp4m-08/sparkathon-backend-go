package controllers

import (
	"net/http"
	"upc-backend-sparkathon/main/models"
	"upc-backend-sparkathon/main/services"

	"github.com/gin-gonic/gin"
)

func ScanAndSaveProduct(c *gin.Context) {
	var request struct {
		UPC string `json:"upc"`
	}

	if err := c.BindJSON(&request); err != nil || request.UPC == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing UPC"})
		return
	}

	product, err := services.FetchProductfromUPC(request.UPC)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := services.DB.Create(product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product saved successfully",
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	var ProductList []models.Product
	if err := services.DB.Find(&ProductList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, ProductList)
}

func GetProduct(c *gin.Context) {
	ean := c.Param("ean")

	var thisproduct models.Product
	if err := services.DB.Where("ean = ?", ean).First(&thisproduct).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, thisproduct)
}
