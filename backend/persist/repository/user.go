package repository

import (
	"backend/persist/db"
	"backend/persist/model"

	"gorm.io/gorm"
)

type User interface {
	FindByNameAndPass(name, password string) *model.User
}

type UserImpl struct {
	db *gorm.DB
}

func NewUser() User {
	return &UserImpl{db: db.GetDB()}
}

func (u *UserImpl) FindByNameAndPass(name, password string) *model.User {
	user := model.User{}
	_ = u.db.Where("name = ? AND pass_word = ?", name, password).Find(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}
