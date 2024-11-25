package database

import (
	"gorm.io/gorm"
)

type FloorModel struct {
	gorm.Model
	X          int    `gorm:""`
	Y          int    `gorm:""`
	Name       string `gorm:"index:uniqueIndex"`
	BuildingID uint   `gorm:""`
	TLLat      string `gorm:"column:tl_lat"`
	TLLng      string `gorm:"column:tl_lng"`
	TRLat      string `gorm:"column:tr_lat"`
	TRLng      string `gorm:"column:tr_lng"`
	BLLat      string `gorm:"column:bl_lat"`
	BLLng      string `gorm:"column:bl_lng"`
	BRLat      string `gorm:"column:br_lat"`
	BRLng      string `gorm:"column:br_lng"`
}

func (FloorModel) TableName() string {
	return "tb_floor"
}
