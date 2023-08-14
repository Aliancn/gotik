package model

import "time"

// 评论
type Comment struct {
	ID         uint      `gorm:"primaryKey"`
	CreateDate time.Time // 8月14日 新增字段
	Content    string

	VideoID uint
	UserID  uint
}
