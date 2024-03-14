package authController

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Issuer string

const (
	AppIssuer    Issuer = "App"
	Oauth2Issuer Issuer = "Oauth2"
)

type JwtClaim struct {
	Iss Issuer `json:"iss"` // issuer : might be app or oauth2
	Id  int    `json:"id"`  // subject
	Exp int64  `json:"exp"` // expiry ts
	Iat int64  `json:"iat"` // issued at ts
	jwt.RegisteredClaims
}

func CreateToken(id int) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenClaims := &JwtClaim{Iss: AppIssuer, Id: id, Exp: time.Now().Add(time.Hour * 24 * 7).Unix(), Iat: time.Now().Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*JwtClaim, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*JwtClaim); ok {
		if checkTokenExpiry(claims.Exp) {
			return claims, nil
		}
	}

	return nil, errors.New("invalid token")
}

func checkTokenExpiry(exp int64) bool {
	now := time.Now().Unix()

	return now < exp
}
