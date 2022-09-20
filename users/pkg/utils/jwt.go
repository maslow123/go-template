package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id     int64
	UserID int32
}

func (w *JwtWrapper) GenerateToken(userID int32) (signedToken string, err error) {
	claims := &jwtClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "invalid-token", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("Couldn't parse claim")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("jwt-expired")
	}

	return claims, nil
}
