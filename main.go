package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	"fsanchez.dev/rest/controllers"
	"fsanchez.dev/rest/models"
)

func main() {
	err := godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	if err != nil && PORT == "" {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	db := models.SetupModels()

	if db != nil {
		fmt.Printf("All Right with Database")
	}

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "RESTful with two Endpoint /users and /notes - v1.0.0"})
	})

	r.GET("/notes", controllers.FindNotes)
	r.POST("/notes", controllers.CreateNote)
	r.GET("/notes/:id", controllers.FindNote)
	r.PATCH("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)

	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run(":" + PORT)
}
