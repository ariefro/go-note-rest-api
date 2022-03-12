package auth

import (
	"errors"

	"github.com/ariefro/notes-server/database"
	"github.com/ariefro/notes-server/model"
	"gorm.io/gorm"
)

func GetUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User

	err := db.Where(&model.User{Email: e}).Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User

	err := db.Where(&model.User{Username: u}).Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}