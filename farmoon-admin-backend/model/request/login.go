package request

import "github.com/golang-jwt/jwt"

type SysJwtClaims struct {
	Username  string `json:"username"`
	RefreshAt int64  `json:"refreshAt"`
	jwt.StandardClaims
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
