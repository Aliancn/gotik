package getuserinfo

import (
	"gotik/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Result struct {
	Code     ResultCode
	UserInfo *model.User
	IsFollow bool
}

type ResultCode int

const (
	ResultCodeOK ResultCode = iota
	ResultCodeTargetNotFound
)

// DoGetUserInfo is not exists, return nil, err
// selfUserID == -1 那么表示未登录用户
func DoGetUserInfo(db *gorm.DB, selfUserID int, targetUserID int) (*Result, error) {
	result := Result{}

	tx := db.Begin()
	defer tx.Rollback()

	// 先锁住selfUser和targetUserID
	err := tx.Clauses(clause.Locking{
		Strength: "SHARE",
	}).Where("id = ? OR id = ?", selfUserID, targetUserID).First(&model.User{}).Error
	if err != nil {
		return nil, err
	}

	var targetUser model.User
	err = db.Where("id = ?", targetUserID).Find(&targetUser).Error
	if err == gorm.ErrRecordNotFound {
		result.Code = ResultCodeTargetNotFound
		return &result, nil
	}

	if selfUserID == -1 {
		result.Code = ResultCodeOK
		result.IsFollow = false
		result.UserInfo = &targetUser
		return &result, nil
	}

	// selfUserID != -1, 那么需要查询下关注信息

	var cnt int64
	err = db.Model(&model.UserFollow{}).Where("user_id = ? AND followed_user_id = ?", selfUserID, targetUserID).Count(&cnt).Error
	if err != nil {
		return nil, err
	}

	result.Code = ResultCodeOK
	result.UserInfo = &targetUser
	switch cnt {
	case 0:
		result.IsFollow = false
	case 1:
		result.IsFollow = true
	default:
		panic("unreachable")
	}

	return &result, err
}
