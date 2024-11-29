package database

import (
	"time"

	"gorm.io/gorm"
)

type InvitationModel struct {
	gorm.Model
	StartTimeStamp   time.Time `json:"start_timestamp" gorm:"type:timestamptz"`
	ExpiredTimeStamp time.Time `json:"expired_timestamp" gorm:"type:timestamptz"`
	IsAccept         string    `json:"is_accept" gorm:"type:char(1)"`
	HomeID           uint      `json:"home_id"`
	HomeFK           UserModel `gorm:"foreignKey:HomeID"`
	AwayID           uint      `json:"away_id"`
	AwayFK           UserModel `gorm:"foreignKey:AwayID"`
	Token            string    `json:"token" gorm:"column:token; index"`
}

func (InvitationModel) TableName() string {
	return "tb_invitation"
}
