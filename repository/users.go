package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) AddUser(user model.User) error {

	res := u.db.Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *UserRepository) UserAvail(cred model.User) error {

	res := u.db.Model(&model.User{}).First(&cred)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (u *UserRepository) CheckPassLength(pass string) bool {
	return len(pass) <= 5
}

func (u *UserRepository) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}
