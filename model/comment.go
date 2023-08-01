package model

// 评论
type Comment struct {
	ID uint `gorm:"primaryKey"`

	Content string

	VideoID uint
	UserID  uint
}
