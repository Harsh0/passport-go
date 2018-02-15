package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/harsh0/passport-go"
	"github.com/subosito/gotenv"
)

var GoogleStrategy passport.Strategy

func init() {
	gotenv.Load()
	GoogleStrategy = passport.GoogleStrategy(map[string]string{
		"clientID":     os.Getenv("GOOGLE_CLIENT_ID"),
		"clientSecret": os.Getenv("GOOGLE_CLIENT_SECRET"),
		"callbackURL":  os.Getenv("HOST") + "/google/callback",
	})
}

/* main will be run during the build process.*/
func main() {
	//for release mode
	// gin.SetMode(gin.ReleaseMode)
	//Initialize router
	r := gin.Default()
	r.GET("/google/signin", func(c *gin.Context) {
		url, err := GoogleStrategy.GetAuthURL("random2")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Some Error Occured, Please try again"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"url": url})
	})
	r.GET("/google/callback", func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		profile, err := GoogleStrategy.Authenticate(code, state)
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
