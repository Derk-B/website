package main

import (
	"api/db"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	projects := db.GetProjects(4)

	for _, p := range projects {
		fmt.Println(p.Title)
	}

	fmt.Println("Initiating REST api on port 8080")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + os.Getenv("API_PORT")) // listen and serve on 0.0.0.0:8080
}
