package main

import (
    //"ecommerce-api/routes"
   "E-COMMERCE API/routes"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    r := gin.Default()
    routes.SetupRoutes(r)
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
