package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harsh0/passport-go"
	"github.com/subosito/gotenv"
)

var GithubStrategy passport.Strategy

func init() {
	gotenv.Load()
	GithubStrategy = passport.GithubStrategy(map[string]string{
		"clientID":     os.Getenv("GITHUB_CLIENT_ID"),
		"clientSecret": os.Getenv("GITHUB_CLIENT_SECRET"),
		"callbackURL":  os.Getenv("HOST") + "/github/callback",
	})
}

/* main will be run during the build process.*/
func main() {
	//for release mode
	// gin.SetMode(gin.ReleaseMode)
	//Initialize router
	r := gin.Default()
	//Use CORS Middleware
	r.GET("/github/signin", func(c *gin.Context) {
		url, err := GithubStrategy.GetAuthURL()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Some Error Occured, Please try again"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"url": url})
	})
	r.GET("/github/callback", func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		profile, err := GithubStrategy.Authenticate(code, state)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Some Error Occured, Please try again"})
			return
		}
		fmt.Println(profile)
		c.JSON(http.StatusOK, gin.H{"message": "successfully logged in"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})
	// run server on port
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
