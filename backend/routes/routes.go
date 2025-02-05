package routes

import (
	"github.com/akhan3247/spotify-music-data/handlers"
	"github.com/akhan3247/spotify-music-data/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// API routes group
	api := router.Group("/api")
	{
		// Auth routes
		api.GET("/login", handlers.Login)
		api.GET("/callback", handlers.Callback)
		api.POST("/logout", handlers.Logout)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// User routes
			protected.GET("/profile", handlers.GetProfile)
			protected.GET("/top-tracks", handlers.GetTopTracks)
			protected.GET("/followed-artists", handlers.GetFollowedArtists)
			protected.GET("/playlists", handlers.GetPlaylists)
			protected.GET("/top-artists", handlers.GetTopArtists)
		}
	}
}
