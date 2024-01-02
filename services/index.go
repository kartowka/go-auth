package services

import (
	"github.com/antfley/go-auth/services/user"
	"gorm.io/gorm"
)

var (
	UserService user.UserService
)

func InjectDBIntoServices(db *gorm.DB) {
	UserService.DB = db
}
