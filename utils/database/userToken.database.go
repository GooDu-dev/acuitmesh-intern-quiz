package database

import (
	"time"
)

type UserTokenModel struct {
	UserID    int       `json:"user_id"`
	UserFK    UserModel `gorm:"foreignKey:UserID"`
	Token     string    `json:"token" gorm:"column:token; index; type:char(32)"`
	OnUsed    string    `json:"on_used" gorm:"column:on_used; type:char(1);"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at; type:timestamptz;"`
}

func (UserTokenModel) TableName() string {
	return "tb_user_token"
}
