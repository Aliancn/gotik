package getcommentlist

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gotik/model"
)

type Result struct {
	Code        ResultCode
	CommentList []model.Comment
}

type ResultCode int

const (
	ResultCodeOK ResultCode = iota
	ResultCodeTargetNotFound
)

func DoGetFavorList(db *gorm.DB, videoID int) (*Result, error) {
	result := Result{}

	tx := db.Begin()
	defer tx.Rollback()

	// 为comment加锁
	err := tx.Clauses(clause.Locking{
		Strength: "SHARE",
	}).Where("video_id = ?", videoID).Find(&model.Comment{}).Error
	if err != nil {
		return nil, err
	}

	// 获取comment的内容
	var comment []model.Comment
	err = tx.Where("video_id = ? ", videoID).Find(&comment).Error
	if err == gorm.ErrRecordNotFound {
		result.Code = ResultCodeTargetNotFound
		return &result, nil
	}

	result.Code = ResultCodeOK
	result.CommentList = comment
	return &result, nil
}
