package main

import (
	"api/routes"
	"fmt"
	"os"
	"strconv"

	"api/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// Setup cors

	// Only allow origins in debug mode, since in release mode
	// the domain of the api will be the same as the website domain.
	isReleaseMode := os.Getenv("MODE") == "release"

	fmt.Println(isReleaseMode)

	var corsConfig cors.Config

	if isReleaseMode {
		corsConfig = cors.DefaultConfig()
	} else {
		corsConfig = cors.DefaultConfig()
		corsConfig.AllowAllOrigins = true
	}

	r.Use(cors.New(corsConfig))

	// Setup db connection
	ctx, dbConn := db.ConnectToDB()
	defer ctx.Done()
	defer dbConn.Close()

	// Define routes
	r.GET("/project", func(c *gin.Context) {
		projectID, err := strconv.Atoi(c.Query("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": "projectID not a number",
			})
			panic(err)
		}

		c.JSON(200, gin.H{
			"message": routes.ReturnProject(projectID, ctx, dbConn),
		})
	})

	r.GET("/projects", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": routes.ReturnProjects(ctx, dbConn),
		})
	})

	r.GET("/blogposts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": routes.ReturnBlogPosts(ctx, dbConn),
		})
	})

	r.Run(":" + os.Getenv("API_PORT"))
}
