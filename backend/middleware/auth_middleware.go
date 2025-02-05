package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	//"github.com/akhan3247/spotify-music-data/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		// Try to get token from cookie
		cookieToken, err := c.Cookie("spotify_token")
		if err == nil && cookieToken != "" {
			token = cookieToken
		} else {
			// Fallback: check for token in the Authorization header
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication token"})
			c.Abort()
			return
		}

		// Create Spotify client with the token
		tok := &oauth2.Token{AccessToken: token}
		httpClient := spotifyauth.New().Client(context.Background(), tok)
		client := spotify.New(httpClient)

		// Optionally you can validate the token here by making a test request

		// Store the client in the context for downstream handlers
		c.Set("spotify", client)
		c.Next()
	}
}
