package model

// 维护user和video的关系, 一个video属于一个用户, 但一个用户可以拥有多个video
type UserVideo struct {
	ID uint `gorm:"primaryKey"`

	UserID  uint
	VideoID uint
}
