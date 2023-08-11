package getuserinfo

import (
	"gotik/model"

	"gorm.io/gorm"
)

// is not exists, return nil, err
func DoGetUserInfo(db *gorm.DB, id uint) (*model.User, error) {
	result := model.User{}
	err := db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func IsAFollowingB(db *gorm.DB, A uint, B uint) (bool, error) {
	cnt := int64(0)
	err := db.Where("user_id = ? AND followed_user_id = ?", A, B).Count(&cnt).Error
	return cnt == 1, err
}
