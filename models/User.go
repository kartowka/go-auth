package models

import (
	"fmt"
	"time"

	"github.com/antfley/go-auth/config"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model   `json:"-"`
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"unique"`
	Password     string `json:"-"`
	RefreshToken string
}

func (u *User) CreateJWT() *jwt.Token {
	var t *jwt.Token
	claims := jwt.RegisteredClaims{
		Issuer:    fmt.Sprint(u.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t
}
func (u *User) CreateRefreshToken() string {
	var (
		t     *jwt.Token
		key   []byte
		token string
	)
	key = []byte(config.Config("JWT_SECRET"))
	claims := jwt.RegisteredClaims{
		Issuer:    fmt.Sprint(u.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(key)
	if err != nil {
		panic(err)
	}
	return token
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Update("refresh_token", u.CreateRefreshToken())
	return nil
}
