package database

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50); uniqueIndex;"`
	Win      uint   `json:"wins" gorm:"default:0"`
	Lose     uint   `json:"loses" gorm:"default:0"`
	Draw     uint   `json:"draws" gorm:"default:0"`
	Total    uint   `json:"total" gorm:"default:0"`
	Mail     string `json:"mail" gorm:"uniqueIndex"`
	Pwd      string `json:"pwd" gorm:"password"`
}

func (UserModel) TableName() string {
	return "tb_user"
}
