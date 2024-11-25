package database

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50); uniqueIndex;"`
	Win      uint   `json:"wins"`
	Lose     uint   `json:"loses"`
	Draw     uint   `json:"draws"`
	Total    uint   `json:"total"`
}

func (UserModel) TableName() string {
	return "tb_user"
}
