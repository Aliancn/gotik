package login

import (
	"errors"
	"gotik/model"
	"gotik/utils/md5"
	"gotik/utils/token"

	"gorm.io/gorm"
)

type Result struct {
	Code   ResultCode
	Token  string
	UserID uint
}

type ResultCode int

const (
	ResultOK ResultCode = iota
	ResultUserNotExists
	ResultPasswordWrong
)

func DoLogin(db *gorm.DB, uname string, pword string) (*Result, error) {
	result := Result{}

	u := model.User{}
	err := db.Where("username = ?", uname).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		result.Code = ResultUserNotExists
		return &result, nil
	}

	if string(u.PasswordMD5) != md5.DoMD5(pword) {
		result.Code = ResultPasswordWrong
		return &result, nil
	}

	result.Code = ResultOK
	result.Token = token.NewToken(int(u.ID), uname)
	result.UserID = u.ID
	return &result, nil
}
