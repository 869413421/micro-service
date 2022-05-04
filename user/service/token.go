package service

import (
	"github.com/869413421/micro-service/user/pkg/model"
	"github.com/869413421/micro-service/user/pkg/repo"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key = []byte("microServiceUserTokenKeySecret")
)

// CustomClaims jwt认证对象
type CustomClaims struct {
	User *model.User
	jwt.StandardClaims
}

// Authble jwt实现接口
type Authble interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *model.User) (string, error)
}

// TokenService token业务对象
type TokenService struct {
	Repo repo.UserRepositoryInterface
}

// NewTokenService token业务初始化
func NewTokenService() Authble {
	return &TokenService{Repo: repo.NewUserRepository()}
}

// Decode 将token字符串转换为token对象
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode 将token对象串转换为token字符串
func (srv *TokenService) Encode(user *model.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "micro.service.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
