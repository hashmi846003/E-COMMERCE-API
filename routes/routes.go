package routes

import (
    "ecommerce-api/controllers"
    "ecommerce-api/utils"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    userRoutes := r.Group("/api/users")
    {
        userRoutes.POST("/register", controllers.Register)
        userRoutes.POST("/login", controllers.Login)
    }
    
    productRoutes := r.Group("/api/products")
    {
        productRoutes.Use(utils.Authenticate())
        productRoutes.POST("/", utils.IsAdmin(), controllers.AddProduct)
        productRoutes.GET("/", controllers.GetProducts)
        productRoutes.POST("/checkout", controllers.Checkout)
    }
}
