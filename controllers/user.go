package controllers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	"github.com/antfley/go-auth/config"
	"github.com/antfley/go-auth/services"
	"github.com/gin-gonic/gin"
)

type UserController struct{}
type ReturnedUser struct {
	ID           uint             `json:"id"`
	Email        string           `json:"email"`
	RefreshToken string           `json:"refresh_token"`
	AccessToken  string           `json:"access_token"`
	ExpiresAt    *jwt.NumericDate `json:"expires_at"`
	Consented_at *jwt.NumericDate `json:"consented_at"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller UserController) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		var response = ErrorResponse{
			Msg: "Validation error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user, err := services.UserService.Login(request.Email, request.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Msg: "Invalid credentials", Err: err})
		return
	}
	isMatchedPassword := services.UserService.DoPasswordsMatch(user.Password, request.Password)
	if !isMatchedPassword {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Msg: "Invalid credentials", Err: err})
		return
	}
	token := user.CreateJWT()
	user.RefreshToken = user.CreateRefreshToken()
	services.UserService.DB.Save(&user)
	expire_at, _ := token.Claims.GetExpirationTime()
	consented_at, _ := token.Claims.GetIssuedAt()
	accessToken, _ := token.SignedString([]byte(config.Config("JWT_SECRET")))
	modifiedUser := ReturnedUser{
		ID:           user.ID,
		Email:        user.Email,
		RefreshToken: user.RefreshToken,
		AccessToken:  accessToken,
		ExpiresAt:    expire_at,
		Consented_at: consented_at,
	}
	c.JSON(http.StatusOK, SuccessResponse{Status: true, Data: modifiedUser})
}
func (controller UserController) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		var response = ErrorResponse{
			Msg: "Validation error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var (
		key          []byte
		t            *jwt.Token
		signedString string
	)
	key = []byte(config.Config("JWT_SECRET"))
	user, err := services.UserService.CreateOne(request.Email, request.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Msg: "Error register!", Err: err})
		return
	}
	t = user.CreateJWT()
	expire_at, _ := t.Claims.GetExpirationTime()
	consented_at, _ := t.Claims.GetIssuedAt()
	signedString, err = t.SignedString(key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Msg: "Error signing token", Err: err})
		return
	}
	modifiedUser := ReturnedUser{
		ID:           user.ID,
		Email:        user.Email,
		RefreshToken: user.RefreshToken,
		AccessToken:  signedString,
		ExpiresAt:    expire_at,
		Consented_at: consented_at,
	}
	c.JSON(http.StatusCreated, SuccessResponse{Status: true, Data: modifiedUser})
}
