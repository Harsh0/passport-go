package passport

import (
	"fmt"
	"goodcop/passport"

	"github.com/gin-gonic/gin"
)

/*GithubSignupForUser allows user signup with github*/
func GithubSignupForUser() {
	passport.GetAuthURL("github", encrypted)
}

/*GoogleSignupForUser allows user signup with google*/
func GoogleSignupForUser() {
	url, err := passport.GetAuthURL("google", encrypted)
}

/*TwitterSignupForUser allows user signup with twitter*/
func TwitterSignupForUser(c *gin.Context) {
	fmt.Println("Twitter signup")
}

/*GithubCallback allows user signup with github*/
func GithubCallback() {
	// profile, err := passport.Authenticate("github", code, state)
}
