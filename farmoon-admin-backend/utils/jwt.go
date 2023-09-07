package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/model/request"
	"strconv"
	"time"
)

func CreateToken(username string) (ss string) {
	//jwtIssuer := GetConfigBackend("jwtIssuer")
	//if jwtIssuer == "" {
	//	return ""
	//}
	jwtIssuer := "Farmoon-Admin"
	//jwtKey := GetConfigBackend("jwtKey")
	//if jwtKey == "" {
	//	return ""
	//}
	jwtKey := "Farmoon-Admin"
	//jwtExpiresAt := GetConfigBackend("jwtExpiresAt")
	//if jwtExpiresAt == "" {
	//	return ""
	//}
	jwtExpiresAt := "14400"
	//jwtRefreshAt := GetConfigBackend("jwtRefreshAt")
	//if jwtRefreshAt == "" {
	//	return ""
	//}
	jwtRefreshAt := "86400"
	jwtExpiresAtInt64, err := strconv.ParseInt(jwtExpiresAt, 10, 64)
	if err != nil {
		return ""
	}
	jwtRefreshAtInt64, err := strconv.ParseInt(jwtRefreshAt, 10, 64)
	if err != nil {
		return ""
	}
	claims := request.SysJwtClaims{
		Username:  username,
		RefreshAt: time.Now().Unix() + jwtRefreshAtInt64,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + jwtExpiresAtInt64,
			Issuer:    jwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString([]byte(jwtKey))
	if err != nil {
		return ""
	}
	return ss
}
