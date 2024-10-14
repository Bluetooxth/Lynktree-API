package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"lynktree/config"
	"lynktree/models"
)

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	collection := config.DB.Collection("users")

	var userUpdate models.UserModel
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"username": userUpdate.Username,
			"name":     userUpdate.Name,
			"email":    userUpdate.Email,
			"links":    userUpdate.Links,
		},
	}

	if userUpdate.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userUpdate.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating password"})
			return
		}
		update["$set"].(bson.M)["password"] = string(hashedPassword)
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}