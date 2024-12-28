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
	authURL := dropbox.DropboxConfig.AuthCodeURL("state-token", oauth2.SetAuthURLParam("token_access_type", "offline"))

	c.JSON(http.StatusOK, gin.H{"authURL": authURL})
}

func HandleDropboxCallback(c *gin.Context) {
	var dropboxToken models.Token

	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No code in the callback request"})
		return
	} 

	token, err := dropbox.DropboxConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token", "details": err.Error()})
		return
	}

	accountID := token.Extra("account_id").(string)
	accessToken := token.AccessToken
	refreshToken := token.RefreshToken
	expiryTime := token.Expiry
	typeToken := token.TokenType
	userID := token.Extra("uid").(string)

	var existingUser models.Token

	if err := database.DB.Where("user_id = ? AND account_id = ?", userID, accountID).First(&existingUser).Error; err != nil{
		if err.Error() == "record not found" {
			dropboxToken = models.Token{
				AccountID: accountID,
				AccessToken: accessToken,
				RefreshToken: refreshToken,
				ExpiryTime: expiryTime,
				TokenType: typeToken,
				UserID: userID,
			}
			if err := database.DB.Create(&dropboxToken).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store token"})
				return
			}
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed", "details": err.Error()})
			return
		}
	}else{
		existingUser.AccessToken = accessToken
		existingUser.RefreshToken = refreshToken
		existingUser.ExpiryTime = expiryTime

		if err := database.DB.Save(&existingUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update token"})
			return
		}
	}

	jwt, err := utils.GenerateJWT(accessToken)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Failed to generate token"})
	}
	c.SetCookie("access_token", jwt, 3600, "/", "localhost", false, false);
	if err := GetCurrentUser(accessToken); err != nil{
		fmt.Println("Error:", err)
	}

	c.Redirect(http.StatusFound, "http://localhost:3000/");
}