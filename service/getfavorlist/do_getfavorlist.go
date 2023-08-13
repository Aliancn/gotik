package getfavorlist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gotik/model"
)

type Result struct {
	Code      ResultCode
	VideoList []model.Video
	VideoNum  int64
}

type ResultCode int

const (
	ResultCodeOK ResultCode = iota
	ResultCodeTargetNotFound
)

func DoGetFavorList(db *gorm.DB, userID int) (*Result, error) {
	result := Result{}

	tx := db.Begin()
	defer tx.Rollback()

	// 先锁住user_favor
	err := tx.Clauses(clause.Locking{
		Strength: "SHARE",
	}).Where("id = ?", userID).First(&model.UserFavorite{}).Error
	if err != nil {
		return nil, err
	}

	// 查询userFavorite表
	// 得到videoIDs
	var userFavor []model.UserFavorite
	err = tx.Where("user_id = ?", userID).Find(&userFavor).Error
	if err == gorm.ErrRecordNotFound {
		result.Code = ResultCodeTargetNotFound
		return &result, nil
	}

	var videoIDs []uint
	for _, favor := range userFavor {
		videoIDs = append(videoIDs, favor.VideoID)
	}

	// 根据videoIDs查询对应的video
	err = tx.Where("id IN (?)", videoIDs).Find(&result.VideoList).Count(&result.VideoNum).Error
	if err == gorm.ErrRecordNotFound {
		result.Code = ResultCodeTargetNotFound
		return &result, nil
	}

	result.Code = ResultCodeOK
	return &result, err
}
