package publish

import (
	"gotik/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Result struct {
	Code ResultCode
}

type ResultCode int

const (
	ResultCodeOK ResultCode = iota
)

func DoPublish(db *gorm.DB, selfUserID int, title string, coverURL string, videoURL string) (*Result, error) {
	result := Result{}

	tx := db.Begin()
	defer tx.Rollback()

	// 先锁住user, 会更新work数目
	user := model.User{}

	err := tx.Clauses(clause.Locking{
		Strength: "UPDATE",
	}).Where("id = ?", selfUserID).First(&user).Error
	if err != nil {
		return nil, err
	}

	err = tx.Create(&model.Video{Title: title, AuthorID: user.ID,
		PlayURL: videoURL, CoverURL: coverURL, FavoriteCount: 0, CommentCount: 0, PublishedAt: time.Now()}).Error
	if err != nil {
		return nil, err
	}

	user.WorkCount++
	err = tx.Save(&user).Error
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	result.Code = ResultCodeOK
	return &result, nil
}
