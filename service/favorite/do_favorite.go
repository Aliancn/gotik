package favorite

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gotik/model"
)

type Result struct {
	Code ResultCode
}

type ResultCode int

const (
	ResultCodeOK ResultCode = iota
	ResultCodeOptionErr
)

func DOFavor(db *gorm.DB, userID int, videoID int) (*Result, error) {
	result := Result{}

	tx := db.Begin()
	defer tx.Rollback()

	// 创建数据
	userFavorite := model.UserFavorite{
		UserID:  uint(userID),
		VideoID: uint(videoID),
	}

	//根据ID修改数据 新增
	err := tx.Create(&userFavorite).Error
	if err != nil {
		result.Code = ResultCodeOptionErr
		return &result, err
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		result.Code = ResultCodeOptionErr
		return &result, err
	}

	// return
	result.Code = ResultCodeOK
	return &result, nil
}

func DOUnFavor(db *gorm.DB, userID int, videoID int) (*Result, error) {
	result := Result{}

	tx := db.Begin()
	defer tx.Rollback()

	//根据ID修改数据 删除
	if err := tx.Clauses(clause.Locking{Strength: "SHARE"}).
		Where("user_id = ?  AND video_id = ? ", userID, videoID).
		Delete(&model.UserFavorite{}).Error; err != nil {
		result.Code = ResultCodeOptionErr
		return &result, err
	}

	// commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		result.Code = ResultCodeOptionErr
		return &result, err
	}

	// return
	result.Code = ResultCodeOK
	return &result, nil
}
