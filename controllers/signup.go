package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"lynktree/config"
	"lynktree/models"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

func Signup(c *gin.Context) {
	var req SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to hash password"})
		return
	}

	collection := config.DB.Collection("users")
	var existingUser models.UserModel
	err = collection.FindOne(context.TODO(), bson.M{"email": req.Email}).Decode(&existingUser)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		return
	}

	user := models.UserModel{
		ID:        primitive.NewObjectID(),
		Username:  req.Username,
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(passwordHash),
		Links:     []models.Links{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = collection.InsertOne(context.TODO(), user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during signup"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User signup successful"})
}
