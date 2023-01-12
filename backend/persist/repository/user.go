package repository

import (
	"backend/persist/db"
	"backend/persist/model"

	"gorm.io/gorm"
)

type User interface {
	FindByNameAndPass(name, password string) *model.User
	CreateUser(user *model.User) (uint, error)
	UpdateUser(user *model.User) error
	DeleteUser(userID uint) error
	GetUser(userID uint) (*model.User, error)
}

type UserImpl struct {
	db *gorm.DB
}

func NewUser() User {
	return &UserImpl{db: db.GetDB()}
}

func (u *UserImpl) FindByNameAndPass(name, password string) *model.User {
	user := model.User{}
	_ = u.db.Where("account = ? AND pass_word = ?", name, password).Find(&user)
	if user.ID == 0 {
		return nil
	}
	return &user
}

func (u *UserImpl) CreateUser(user *model.User) (uint, error) {
	tx := u.db.Create(user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return user.ID, nil
}

func (u *UserImpl) UpdateUser(user *model.User) error {
	tx := u.db.Updates(user)
	return tx.Error
}

func (u *UserImpl) DeleteUser(userID uint) error {
	tx := u.db.Delete(&model.User{
		Model: gorm.Model{
			ID: userID,
		},
	})
	return tx.Error
}

func (u *UserImpl) GetUser(userID uint) (*model.User, error) {
	user := model.User{}
	user.ID = userID
	tx := u.db.First(&user)
	return &user, tx.Error
}
