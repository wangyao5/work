package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const accessTokenSignedString string = "secret"
const refreshTokenSignedString string = "secret"

// Claims custom token
type Claims struct {
	API  string `json:"api,omitempty"`  // API接口
	Menu string `json:"menu,omitempty"` // 菜单权限
	jwt.StandardClaims
}

//CreateAccessToken 创建自定义Token.
func CreateAccessToken(claims Claims) (string, error) {
	claims.ExpiresAt = time.Now().Add(time.Minute * 60).Unix()
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimToken.SignedString([]byte(accessTokenSignedString))
}

//CreateRefreshToken 创建refresh Token.
func CreateRefreshToken() (string, error) {
	claims := Claims{}
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimToken.SignedString([]byte(refreshTokenSignedString))
}
