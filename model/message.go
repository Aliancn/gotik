package model

import "time"

type Message struct {
	ID uint `gorm:"primaryKey"`

	SenderUserID   uint
	ReceiverUserID uint
	Content        string

	CreatedAt time.Time
}
