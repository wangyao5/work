package security

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

const accessTokenSignedString string = "secret"
const refreshTokenSignedString string = "secret"

// Claims custom token
type Claims struct {
	UID string `json:"uid,omitempty"` //UID
	API string `json:"api,omitempty"` // API接口
	jwt.StandardClaims
}

//CreateAccessToken 创建自定义Token.
func CreateAccessToken(claims Claims, expiresAt int64) (string, error) {
	claims.ExpiresAt = expiresAt
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimToken.SignedString([]byte(accessTokenSignedString))
}

//CreateRefreshToken 创建refresh Token.
func CreateRefreshToken(claims Claims, expiresAt int64) (string, error) {
	claims.ExpiresAt = expiresAt
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return claimToken.SignedString([]byte(refreshTokenSignedString))
}

//ValidateAccessToken 验证Token是否有效
func ValidateAccessToken(accessToken string) (claims *Claims, ok bool) {
	token, err := jwt.ParseWithClaims(accessToken, new(Claims),
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte(accessTokenSignedString), nil
		})
	if err != nil {
		log.Fatal(err)
	}
	claims, ok = token.Claims.(*Claims)
	if !token.Valid {
		ok = false
	}
	return
}

//ValidateRefreshToken 验证Token是否有效
func ValidateRefreshToken(refreshToken string) (claims *Claims, ok bool) {
	token, err := jwt.ParseWithClaims(refreshToken, new(Claims),
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte(refreshTokenSignedString), nil
		})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok = token.Claims.(*Claims)
	if !token.Valid {
		ok = false
	}
	return
}
