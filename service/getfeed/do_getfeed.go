package getfeed

import (
	"gotik/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ResultDataEntry struct {
	Video           *model.Video
	Author          *model.User
	IsFollowAuthor  bool
	IsFavoriteVideo bool
}

type Result struct {
	Code ResultCode
	List []*ResultDataEntry
}

type ResultCode int

const (
	ResultCodeOK ResultCode = iota
)

// TODO: 不一致问题, 先查entry再查user的问题
//
//	follow和followed_count
func DoGetFeed(db *gorm.DB, selfUserID int, upperID int) (*Result, error) {
	resultEntries := make([]*ResultDataEntry, 0, 10)

	var videos []*model.Video
	err := db.Where("id < ?", upperID).Clauses(clause.Locking{
		Strength: "SHARE",
	}).Order("id DESC").Limit(10).Find(&videos).Error
	if err != nil {
		return nil, err
	}

	for _, v := range videos {
		entry := ResultDataEntry{}
		user := model.User{}
		err := db.Where("id = ?", v.AuthorID).First(&user).Error
		if err != nil {
			return nil, err
		}
		if selfUserID == -1 {
			entry.IsFollowAuthor = false
			entry.IsFavoriteVideo = false
		} else {
			var cnt int64
			err := db.Model(&model.User{}).Where("user_id = ? AND followed_user_id = ?", selfUserID, user.ID).Count(&cnt).Error
			if err != nil {
				return nil, err
			}
			entry.IsFollowAuthor = (cnt == 1)

			err = db.Model(&model.UserFavorite{}).Where("user_id = ? AND video_id = ?", selfUserID, user.ID).Count(&cnt).Error
			if err != nil {
				return nil, err
			}
			entry.IsFavoriteVideo = (cnt == 1)
		}
		entry.Author = &user
		entry.Video = v
		resultEntries = append(resultEntries, &entry)
	}

	return &Result{
		Code: ResultCodeOK,
		List: resultEntries,
	}, nil
}
