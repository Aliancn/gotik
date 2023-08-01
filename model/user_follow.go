package model

// 维护用户间的关注信息
type UserFollow struct {
	ID uint `gorm:"primaryKey"`

	UserID uint

	// 被Follow的User的ID
	FollowedUserID uint
}
