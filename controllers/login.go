package controllers

import (
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "golang.org/x/crypto/bcrypt"
    "lynktree/config"
    "lynktree/models"
    "lynktree/utils"
)

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6,max=100"`
}

func Login(c *gin.Context) {
    var req LoginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    collection := config.DB.Collection("users")
    var user models.UserModel

    err := collection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := utils.GenerateJWT(user.Username, user.ID.Hex())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    http.SetCookie(c.Writer, &http.Cookie{
        Name:   "token",
        Value:  token,
        MaxAge: 3600,
        Path:   "/",
    })

    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
        "user": gin.H{
            "username": user.Username,
        },
    })
}