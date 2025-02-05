package handlers

import (
	"context"
	"net/http"

	"github.com/akhan3247/spotify-music-data/models"
	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

func GetProfile(c *gin.Context) {
	client := c.MustGet("spotify").(*spotify.Client)

	spotifyUser, err := client.CurrentUser(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch profile"})
		return
	}

	images := make([]models.Image, len(spotifyUser.Images))
	for i, image := range spotifyUser.Images {
		images[i] = models.Image{
			URL:    image.URL,
			Width:  int(image.Width),
			Height: int(image.Height),
		}
	}
	user := models.UserProfile{
		DisplayName: spotifyUser.DisplayName,
		Followers:   int(spotifyUser.Followers.Count),
		Images:      images,
	}

	c.JSON(http.StatusOK, user)
}
