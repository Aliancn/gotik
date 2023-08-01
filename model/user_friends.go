package model

// 保存用户间的好友关系
// 存这张表的时候要注意, ID较小的给LowerUserID
type UserFriends struct {
	ID uint `gorm:"primaryKey"`

	LowerUserID uint
	UpperUserID uint
}
