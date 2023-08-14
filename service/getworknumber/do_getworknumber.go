package getworknumber

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gotik/model"
)

type Result struct {
	Code      ResultCode
	WorkCount uint
}
type ResultCode int

const (
	ResultCodeOK ResultCode = iota
	ResultCodeTargetNotFound
)

// DOGetWorkNumber  根据userID获取用户作品数量
func DOGetWorkNumber(db *gorm.DB, userID uint) (*Result, error) {
	result := Result{}

	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Clauses(clause.Locking{
		Strength: "SHARE",
	}).Where("id = ? ", userID).Find(&model.UserVideo{}).Error
	if err != nil {
		return nil, err
	}
	var workCount int64
	err = tx.Model(&model.UserVideo{}).Where("id = ?", userID).Count(&workCount).Error
	if err == gorm.ErrRecordNotFound {
		result.Code = ResultCodeTargetNotFound
		return &result, nil
	}

	result.Code = ResultCodeOK
	result.WorkCount = uint(workCount)
	return &result, nil

}
