package main

import (
	"api/routes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"slices"
	"strconv"

	"api/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func IPWhiteList(whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if !slices.Contains(whitelist, ip) {
			c.IndentedJSON(http.StatusForbidden, gin.H{
				"message": "You are not authorised to use this endpoint",
			})
			return
		}
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	// Setup cors

	// Only allow origins in debug mode, since in release mode
	// the domain of the api will be the same as the website domain.
	isReleaseMode := os.Getenv("MODE") == "release"

	var corsConfig cors.Config

	if isReleaseMode {
		corsConfig = cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{"http://localhost:5173"}
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
	r.GET("/blog", func(c *gin.Context) {
		blogID, err := strconv.Atoi(c.Query("id"))

		if err != nil {
			c.JSON(400, gin.H{
				"message": "blogID not a number",
			})
			panic(err)
		}

		c.JSON(200, gin.H{
			"message": routes.ReturnBlogPost(blogID, ctx, dbConn),
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

	whitelist := []string{}
	whitelistParseErr := json.Unmarshal([]byte(os.Getenv("IP_WHITELIST")), &whitelist)

	if whitelistParseErr != nil {
		panic(whitelistParseErr)
	}

	protectedEndpoint := r.Group("/admin")
	protectedEndpoint.Use(IPWhiteList(whitelist))

	protectedEndpoint.POST("/createproject", func(c *gin.Context) {
		var project = db.Project{}
		data, requestBodyErr := io.ReadAll(c.Request.Body)

		if requestBodyErr != nil {
			panic(requestBodyErr)
		}

		json.Unmarshal(data, &project)

		db.InsertProject(project, ctx, dbConn)

		c.JSON(200, gin.H{
			"message": "Project uploaded",
		})
	})

	protectedEndpoint.POST("/createblog", func(c *gin.Context) {
		var blogPost = db.BlogPost{}
		data, requestBodyErr := io.ReadAll(c.Request.Body)

		if requestBodyErr != nil {
			panic(requestBodyErr)
		}

		json.Unmarshal(data, &blogPost)

		db.InsertBlogPost(blogPost, ctx, dbConn)

		c.JSON(200, gin.H{
			"message": "Blogpost uploaded",
		})
	})

	r.Run(":" + os.Getenv("API_PORT"))
}
