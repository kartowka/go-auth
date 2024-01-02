package user

import (
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}
