package controllers

import (
	"net/http"

	"fsanchez.dev/rest/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CreateNoteInput validate input for create note
type CreateNoteInput struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdateNoteInput validate input for update note
type UpdateNoteInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Conten string `json:"content"`
}

// FindNotes Get all notes
func FindNotes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var notes []models.Note
	db.Find(&notes)
	c.JSON(http.StatusOK, gin.H{"data": notes})
}

// FindNote Get a specific note
func FindNote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var note models.Note
	if err := db.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": note})
}

// CreateNote Insert in DB a new note
func CreateNote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note := models.Note{Title: input.Title, Author: input.Author, Content: input.Content}
	db.Create(&note)

	c.JSON(http.StatusOK, gin.H{"data": note})
}

// UpdateNote change content of the note
func UpdateNote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var note models.Note
	if err := db.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	var input UpdateNoteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&note).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": note})
}

// DeleteNote remove from database
func DeleteNote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var note models.Note
	if err := db.Where("id = ?", c.Param("id")).First(&note).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	db.Delete(&note)

	c.JSON(http.StatusOK, gin.H{"data": "Note is delete"})
}
