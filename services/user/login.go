package user

import (
	"github.com/antfley/go-auth/models"
)

func (s *UserService) Login(email, password string) (models.User, error) {
	var user models.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
