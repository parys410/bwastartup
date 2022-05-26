package user

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if(err != nil) {
		return user, err
	} else {
		fmt.Println("UserRepository/Save() is triggered")
		return user, nil
	}
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Table("users").Where("email = ?", email).Find(&user).Error
	if(err != nil) {
		return user, err
	} else if user.ID == 0 {
		newError := errors.New("no user found on that email")
		return user, newError
	} else {
		fmt.Println("UserRepository/FindByEmail() is triggered")
		return user, nil
	}
}