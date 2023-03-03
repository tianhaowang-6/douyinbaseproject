package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 2
	JWTSecret           = "wth"
)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
