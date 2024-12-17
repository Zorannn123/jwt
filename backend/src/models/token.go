package models

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	AccountID 	   string  	 	`json:"account_id"`
	AccessToken    string 	 	`json:"access_token"`
	RefreshToken   string 	 	`json:"refresh_token"`
	ExpiryTime     time.Time 	`json:"expiry_time"`
	TokenType 	   string 	 	`json:"token_type"`
	UserID 		   string 	 	`json:"user_id"`
}