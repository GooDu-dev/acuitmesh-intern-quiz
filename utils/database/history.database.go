package database

import "gorm.io/gorm"

type HistoryModel struct {
	gorm.Model
	Winner       string          `json:"winner" gorm:"type:char(1)"`
	InviteID     uint            `json:"invite_id"`
	InvitationFK InvitationModel `gorm:"foreignKey:InviteID"`
	HomeID       uint            `json:"home_id"`
	HomeFK       UserModel       `gorm:"foreignKey:HomeID"`
	AwayID       uint            `json:"away_id"`
	AwayFK       UserModel       `gorm:"foreignKey:AwayID"`
}

func (HistoryModel) TableName() string {
	return "tb_history"
}
