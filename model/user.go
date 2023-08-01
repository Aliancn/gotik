package model

import "time"

// 用户信息
type User struct {
	ID uint `gorm:"primaryKey"`

	// 用户名
	Username string `gorm:"unique"`
	// 密码的md5摘要
	PasswordMD5 []byte

	// 头像图片
	AvatarURL string
	// 主页大图
	BackgroundImageURL string

	// 获赞数量
	FavoritedCount uint
	// 作品数量
	WorkCount uint
	// 赞的数量
	FavoriteCount uint
	// 关注总数
	FollowCount uint
	// 被关注的数目
	FollowerCount uint

	// 自己是否关注了此用户
	IsFollow bool

	// 个性签名
	Signature string

	// 创建用户的时间
	CreatedAt time.Time
}
