package controllers

import (
    "ecommerce-api/models"
    "ecommerce-api/utils"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
    "net/http"
    "time"
)

var userCollection *mongo.Collection = utils.GetCollection("users")

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    user.HashPassword()
    
    _, err := userCollection.InsertOne(c, user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error registering user"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var input models.User
    var user models.User
    
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    err := userCollection.FindOne(c, bson.M{"email": input.Email}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    
    if !user.CheckPassword(input.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    
    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"token": token})
}
