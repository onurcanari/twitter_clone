package server

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// GetClaimsToken Gets claims and token from request
func GetClaimsToken(r *http.Request) (claims *JwtClaims, token *jwt.Token, err error) {
	c, err := r.Cookie("x-csrf-token")
	if err != nil {
		return
	}
	tokenString := c.Value
	claims = &JwtClaims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return
}

// GetNameFromToken Gets username from request
func GetNameFromToken(r *http.Request) string {
	claims, _, _ := GetClaimsToken(r)
	return claims.Username
}
