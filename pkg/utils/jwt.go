package utils

import (
	"errors"

	"fmt"
	"log"
	"time"

	"github.com/axadjonovsardorbek/tender/config"
	"github.com/golang-jwt/jwt"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateJWTToken(userID string, role string) *Tokens {
	cnf := config.Load()
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(72 * time.Hour).Unix() // Token expires in 3 days
	access, err := accessToken.SignedString([]byte(cnf.JWTSecret))
	if err != nil {
		log.Fatal("error while generating access token : ", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["role"] = role
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(720 * time.Hour).Unix() // Refresh token expires in 30 days
	refresh, err := refreshToken.SignedString([]byte(cnf.JWTSecret))
	if err != nil {
		log.Fatal("error while generating refresh token : ", err)
	}

	return &Tokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	cnf := config.Load()
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(cnf.JWTSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}
	fmt.Print(token.Claims)
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
