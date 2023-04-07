package router

import (
	"challenge294/controllers"
	"challenge294/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication()) // Punya token yang valid ? lanjut : keluar

		productRouter.POST("/", controllers.CreateProduct) // Create, produk yang dibuat akan langsung diset dengan id user yang telah terautentikasi
		// productRouter.GET("/", controllers.ReadProduct) // Read All
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.ReadProduct) // Read Product By ID
		productRouter.PUT("/:productId", middlewares.AdminOnlyAuthorization(), controllers.UpdateProduct) // Update
		productRouter.DELETE("/:productId", middlewares.AdminOnlyAuthorization(), controllers.DeleteProduct) // Delete
	}

	return r
}