package handlers

import (
	"context"
	"net/http"

	"github.com/akhan3247/spotify-music-data/models"
	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

func GetTopTracks(c *gin.Context) {
	client := c.MustGet("spotify").(*spotify.Client)

	spotifyTracks, err := client.CurrentUsersTopTracks(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top tracks"})
		return
	}

	tracks := make([]models.Track, len(spotifyTracks.Tracks))

	for i, track := range spotifyTracks.Tracks {
		images := make([]models.Image, len(track.Album.Images))
		for j, image := range track.Album.Images {
			images[j] = models.Image{
				URL:    image.URL,
				Width:  int(image.Width),
				Height: int(image.Height),
			}
		}

		artists := make([]struct {
			Name string `json:"name"`
		}, len(track.Artists))
		for j, artist := range track.Artists {
			artists[j].Name = artist.Name
		}

		tracks[i] = models.Track{
			Album: struct {
				Images []models.Image `json:"images"`
			}{
				Images: images,
			},
			Name:     track.Name,
			Artists:  artists,
			Duration: int(track.Duration),
			ExternalURLs: struct {
				Spotify string `json:"spotify"`
			}{
				Spotify: track.ExternalURLs["spotify"],
			},
		}
	}

	response := models.TopTracks{
		Items: tracks,
	}

	c.JSON(http.StatusOK, response)
}

func GetFollowedArtists(c *gin.Context) {
	client := c.MustGet("spotify").(*spotify.Client)

	artists, err := client.CurrentUsersFollowedArtists(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch followed artists"})
		return
	}

	c.JSON(http.StatusOK, artists)
}

func GetPlaylists(c *gin.Context) {
	client := c.MustGet("spotify").(*spotify.Client)

	playlists, err := client.CurrentUsersPlaylists(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch playlists"})
		return
	}

	c.JSON(http.StatusOK, playlists)
}

func GetTopArtists(c *gin.Context) {
	client := c.MustGet("spotify").(*spotify.Client)

	spotifyArtists, err := client.CurrentUsersTopArtists(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch top artists"})
		return
	}

	artists := make([]models.Artist, len(spotifyArtists.Artists))

	for i, artist := range spotifyArtists.Artists {
		images := make([]models.Image, len(artist.Images))
		for j, image := range artist.Images {
			images[j] = models.Image{
				URL:    image.URL,
				Width:  int(image.Width),
				Height: int(image.Height),
			}
		}
		artists[i] = models.Artist{
			Name:   artist.Name,
			Images: images,
			ExternalURLs: struct {
				Spotify string `json:"spotify"`
			}{
				Spotify: artist.ExternalURLs["spotify"],
			},
			Genres: artist.Genres,
		}
	}

	response := models.TopArtists{
		Items: artists,
	}

	c.JSON(http.StatusOK, response)
}
