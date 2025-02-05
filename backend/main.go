package main

import (
	"log"
	"os"

	"github.com/akhan3247/spotify-music-data/config"
	"github.com/akhan3247/spotify-music-data/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	log.Printf("SPOTIFY_CLIENT_ID: %s", os.Getenv("SPOTIFY_CLIENT_ID"))
	log.Printf("REDIRECT_URI: %s", os.Getenv("REDIRECT_URI"))
	// Initialize config
	config.InitConfig()

	// Set up Gin
	app := gin.Default()

	// Setup CORS
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Initialize routes
	routes.SetupRoutes(app)

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := app.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
