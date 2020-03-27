package server

import (
	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Password string `xorm:"text" json:"Password"`
	Username string `xorm:"varchar(20) notnull pk unique" json:"Username"`
}

type jwtClaims struct {
	Username string `json:"Username"`
	jwt.StandardClaims
}
