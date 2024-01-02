package user

import (
	"github.com/antfley/go-auth/models"
)

func (s *UserService) CreateOne(email string, password string) (models.User, error) {
	user := models.User{
		Email:    email,
		Password: password,
	}
	passwd, err := s.HashPassword(password)
	if err != nil {
		return user, err
	}
	user.Password = passwd
	if err := s.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
