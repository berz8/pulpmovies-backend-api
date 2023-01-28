package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/berz8/pulpmovies-backend-api/config"
	"github.com/berz8/pulpmovies-backend-api/controllers"
	"github.com/gin-contrib/gzip"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
  // Load env file
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal("Error: failed to load env file")
  }

  if os.Getenv("ENV") == "PRODUCTION" {
    gin.SetMode(gin.ReleaseMode)
  }

  // Start gin server
  r := gin.Default()

  config.ConnectDatabase()
  config.RedisClient()

  r.Use(gzip.Gzip(gzip.DefaultCompression))
  r.Use(CORSMiddleware())

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "name": "Pulpmovies APIs",
      "version": os.Getenv("API_VERSION"),
    })
  })

  r.NoRoute(func(c *gin.Context) {
    c.JSON(http.StatusNotFound, gin.H{
      "name": "Pulpmovies APIs",
      "version": os.Getenv("API_VERSION"),
      "message": "Not found",
    })
  })

  v1 := r.Group("v1")
  {

    auth := new(controllers.AuthController)
    v1.POST("/auth/login", auth.Login)

    user := new(controllers.UserController)
    v1.POST("/user/register", user.Register)
  }

  port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("API_VERSION"))

  r.Run(":" + port)

}
