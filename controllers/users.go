package controllers

import (
	"net/http"

	"fsanchez.dev/rest/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateUserInput validate input data for create
type CreateUserInput struct {
	FirstName string `json:"firstname" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

// UpdateUserInput validate data for update
type UpdateUserInput struct {
	FirstName string `json:"firstname"`
	Email     string `json:"email"`
}

// FindUsers Get all users
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []models.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// FindUser Get a specific user
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser Insert in DB a new user
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{FirstName: input.FirstName, Email: input.Email}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser change content of the user
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser remove from database
func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": "User is delete"})
}
