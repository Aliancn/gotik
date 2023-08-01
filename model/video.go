package model

import "time"

// 作品信息
type Video struct {
	ID uint `gorm:"primaryKey"`

	// 作品的标题
	Title string
	// 发布者的ID
	AuthorID uint

	// 播放地址
	PlayURL string
	// 封面地址
	CoverURL string
	// 获得赞的数目
	FavoriteCount uint
	// 评论数目
	CommentCount uint

	// 自己是否赞了此作品
	IsFavorite bool

	// 发布日期
	PublishedAt time.Time
}
