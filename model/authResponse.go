package model

import "github.com/dgrijalva/jwt-go"

type LoginResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	SignatureKey string `json:"signature_key"`
	Username     string `json:"username"`
	jwt.StandardClaims
}
