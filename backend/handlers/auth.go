package handlers

import (
	"context"
	"net/http"

	"github.com/akhan3247/spotify-music-data/config"

	"github.com/gin-gonic/gin"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURI = "http://localhost:8080/api/callback"

func Login(c *gin.Context) {
	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithClientID(config.Config.SpotifyClientID),
		spotifyauth.WithClientSecret(config.Config.SpotifyClientSecret),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserTopRead,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopePlaylistReadPrivate,
		),
	)

	state := "some-random-state-key"
	url := auth.AuthURL(state)

	// Return detailed debug information
	c.JSON(http.StatusOK, gin.H{
		"url": url,
	})
}

func Callback(c *gin.Context) {
	auth := spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithClientID(config.Config.SpotifyClientID),
		spotifyauth.WithClientSecret(config.Config.SpotifyClientSecret),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserTopRead,
			spotifyauth.ScopeUserFollowRead,
			spotifyauth.ScopePlaylistReadPrivate,
		),
	)

	code := c.Query("code")
	//state := c.Query("state")

	token, err := auth.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token: " + err.Error()})
		return
	}

	// Store token in session/cookie
	c.SetCookie("spotify_token", token.AccessToken, 3600, "/", "", false, false)

	// Return success response instead of redirect
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/dashboard")
}

func Logout(c *gin.Context) {
	c.SetCookie("spotify_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:3000",
	})

}
