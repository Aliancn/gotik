package model

// 关系表, 用来维护点赞信息
type UserFavorite struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	VideoID uint
}
