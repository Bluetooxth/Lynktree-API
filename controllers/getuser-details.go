package controllers

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "lynktree/config"
    "lynktree/models"
    "lynktree/utils"
)

func GetUserDetails(c *gin.Context) {
    tokenString, err := c.Cookie("token")
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in cookies"})
        return
    }

    token, claims, err := utils.ValidateJWT(tokenString)
    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }

    userID := claims.ID
    collection := config.DB.Collection("users")
    var user models.UserModel

    objectID, err := primitive.ObjectIDFromHex(userID)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
        return
    }

    err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}