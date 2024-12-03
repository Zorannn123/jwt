package handlers

import (
	"AUTH/src/database"
	"AUTH/src/models"

	"gorm.io/gorm"
)

func findLoginUser(username string, password string) (*models.User, error) {
	var user models.User

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil{
		return nil, nil
	}

	return &user, nil
}

func checkUsername(username string) (*models.User, error){
	var user models.User

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
