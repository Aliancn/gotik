package register

import (
	"errors"
	"gotik/model"
	"gotik/utils/consts"
	"gotik/utils/md5"
	"time"

	"gorm.io/gorm"
)

type Result struct {
	Code ResultCode
	ID   uint
}

type ResultCode int

const (
	ResultOK ResultCode = iota
	ResultUsernameOccupied
)

func DoRegister(db *gorm.DB, uname string, pword string) (*Result, error) {
	// 创建事务
	tx := db.Begin()

	// 执行插入
	newUser := model.User{
		Username:           uname,
		PasswordMD5:        []byte(md5.DoMD5(pword)),
		AvatarURL:          consts.DefaultAvatar,
		BackgroundImageURL: consts.DefaultBackground,
		FavoritedCount:     0,
		WorkCount:          0,
		FavoriteCount:      0,
		FollowCount:        0,
		FollowerCount:      0,
		Signature:          "默认签名",
		CreatedAt:          time.Now(),
	}

	// 这里用到了个技巧, 利用unique的错误代码来判断用户名是否被注册
	if err := tx.Create(&newUser).Error; err != nil {
		tx.Rollback()

		// 检查是否重复键, 来判断用户名是否重复
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return &Result{
				Code: ResultUsernameOccupied,
			}, nil
		}

		// 可能是网络错误, 这里错误无法处理, 报告给调用者
		return nil, err
	}

	// commit也可能因为网络错误而出错, 需要判断下
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// 返回插入的数据
	return &Result{
		Code: ResultOK,
		ID:   newUser.ID,
	}, nil
}
