package handlers

import (
	"AUTH/src/database"
	"AUTH/src/dropbox"
	"AUTH/src/models"
	"AUTH/src/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func Login(c *gin.Context) {
	var loginReq models.User
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	user, err := findLoginUser(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	existingUser, err := checkUsername(user.Username)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to check username"})
		return
	}

	if existingUser != nil{
		c.JSON(http.StatusConflict, gin.H{"error" : "Username already exists"})
		return
	}

	if err := user.HashPassword(user.Password); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to hash password"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func HandleDropboxLogin(c *gin.Context) {
	authURL := dropbox.DropboxConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	c.JSON(http.StatusOK, gin.H{"authURL": authURL})
	fmt.Println("Auth URL: ", authURL)
}

func HandleDropboxCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	fmt.Println("code ", code)
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No code in the callback request"})
		return
	} 

	token, err := dropbox.DropboxConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token", "details": err.Error()})
		return
	}
	fmt.Println("token ", token.AccessToken)
	fmt.Println("url ", dropbox.DropboxConfig.RedirectURL)

	// c.SetCookie("access_token", token.AccessToken, 3600, "/", "localhost", false, true);

	//c.Redirect(http.StatusFound, "http://localhost:3000/");
	//TODO: store in database and check refresh for access token
	c.JSON(http.StatusOK, gin.H{
		"access_token": token.AccessToken,
	})

}