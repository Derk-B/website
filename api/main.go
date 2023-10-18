package main

import (
	"api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../database/.env")

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/projects", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": routes.ReturnProjects(),
		})
	})

	r.GET("/blogposts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": routes.ReturnBlogPosts(),
		})
	})

	r.Run(":" + os.Getenv("API_PORT"))
}
