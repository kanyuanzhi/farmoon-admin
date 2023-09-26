package utils

import (
	"github.com/kanyuanzhi/farmoon-admin/farmoon-admin-backend/global"
	"golang.org/x/crypto/bcrypt"
)

func EncodeBcrypt(str string) (hs string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		global.FXLogger.Error("encode bcrypt error" + err.Error())
		return "", err
	}
	return string(hash), nil
}

func CompareBcrypt(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
