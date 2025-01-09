package controllers

import (
    "ecommerce-api/models"
    "ecommerce-api/services"
    "ecommerce-api/utils"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
    "time"
)

var productCollection *mongo.Collection = utils.GetCollection("products")

func AddProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product.CreatedAt = time.Now()
    product.UpdatedAt = time.Now()
    
    _, err := productCollection.InsertOne(c, product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding product"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Product added successfully"})
}

func GetProducts(c *gin.Context) {
    cur, err := productCollection.Find(c, bson.D{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
        return
    }
    defer cur.Close(c)
    
    var products []models.Product
    for cur.Next(c) {
        var product models.Product
        cur.Decode(&product)
        products = append(products, product)
    }
    c.JSON(http.StatusOK, products)
}

func Checkout(c *gin.Context) {
    var cart struct {
        Amount float64 `json:"amount"`
    }
    if err := c.ShouldBindJSON(&cart); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create a channel to handle payment status
    done := make(chan bool)
    services.ProcessPayment(cart.Amount, done)

    // Wait for the payment to be processed
    if <-done {
        c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment processing failed"})
    }
}
