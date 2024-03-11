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
	iss Issuer // issuer : might be app or oauth2
	id  int    // subject
	exp int64  // expiry ts
	iat int64  // issued at ts
	jwt.RegisteredClaims
}

func CreateToken(id int) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenClaims := JwtClaim{iss: AppIssuer, id: id, exp: time.Now().Add(time.Hour * 24 * 7).Unix(), iat: time.Now().Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": tokenClaims.iss,
		"id":  tokenClaims.id,
		"exp": tokenClaims.exp,
		"iat": tokenClaims.iat,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*JwtClaim, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*JwtClaim); ok {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
