package utils

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "os"
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "log"
    "github.com/joho/godotenv"
    "context"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Authenticate() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
            c.Abort()
            return
        }
        tokenString = tokenString[len("Bearer "):]
        
        claims := &jwt.StandardClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey,[_{{{CITATION{{{_1{](https://github.com/vicent-dev/mailer/tree/27bb7da8969622366802cd7915513fb666a35bdc/main.go)[_{{{CITATION{{{_2{](https://github.com/shoeguard/main-backend/tree/c62e1fd4067c31f685c8f92b5afa2f05b9d375c7/src%2Fcmd%2Fserver%2Fmodels%2Fuser.go)[_{{{CITATION{{{_3{](https://github.com/Ekenzy-101/Go-Instagram-API/tree/9d3937b28b58d09bf64dd85925beb69617e938e2/handlers%2Fpost.go)[_{{{CITATION{{{_4{](https://github.com/keithweaver/go-boilerplate/tree/759f7ccc84c726c6b491bb3fbdce38b364a8daa5/cars%2Frepository.go)[_{{{CITATION{{{_5{](https://github.com/Raytrox/Coursechecker/tree/8cc23304f4ffbffe712a11b363808cf13a1b7ede/monitor.go)[_{{{CITATION{{{_6{](https://github.com/KisLupin/go-mongodb/tree/d8c657016e6e2e216235ffa6516b13da58f6786b/main%2Fmain.go)[_{{{CITATION{{{_7{](https://github.com/aliparlakci/country-roads/tree/ed7cfc6f32ce84e46377e7cb3d11d964c35996f0/server%2Fcontrollers%2Fcontact_controller.go)