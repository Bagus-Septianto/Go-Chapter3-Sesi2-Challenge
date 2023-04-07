package controllers

import (
	"challenge294/database"
	"challenge294/helpers"
	"challenge294/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims) // ambil userData dari jwt
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID // Product.UserID yang akan dibuat diset dengan ID user yang telah login/ID yang tersimpan di userData/jwt
	
	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func ReadProduct(c *gin.Context) {
	// TODO cek ini lagi
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims) // ambil userData dari jwt
	contentType := helpers.GetContentType(c)

	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))
	// role := userData["role"]

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID // Product.UserID yang akan dibuat diset dengan ID user yang telah login/ID yang tersimpan di userData/jwt
	
	err := db.Debug().First(&Product, "id = ?", productId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// if Product.UserID != userID && role != "admin" { // cek kalau produk yang dicari memiliki userID yang berbeda dengan user yang terautentikasi
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error":   "Unauthorized",
	// 		"message": "you are not allowed to access this data",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, Product)
}

func UpdateProduct(c *gin.Context)  {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims) // ambil userData dari jwt
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func DeleteProduct(c *gin.Context)  {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims) // ambil userData dari jwt
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Where("id = ?", productId).Delete(&models.Product{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}