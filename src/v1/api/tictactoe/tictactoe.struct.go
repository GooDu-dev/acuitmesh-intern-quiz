package tictactoe

import (
	"time"
)

type MatchInfo struct {
	HomePlayer UserMatchCard `json:"home"`
	AwayPlayer UserMatchCard `json:"away"`
	Board      BoardInfo     `json:"board"`
	ExpiredAt  time.Time     `json:"expired_at"`
	IsEnd      string        `json:"is_end"`
	InviteID   int           `json:"invite_id"`
}

type BoardInfo struct {
	Data  []int  `json:"data"`
	Token string `json:"token"`
}

type UserMatchCreateStruct struct {
	HomeID    int        `json:"home_id" gorm:"column:home_id"`
	AwayID    int        `json:"away_id" gorm:"column:away_id"`
	InviteID  int        `json:"invite_id" gorm:"column:invite_id"`
	Token     string     `json:"token" gorm:"column:token"`
	IsEnd     string     `json:"is_end" gorm:"column:is_end"`
	ExpiredAt time.Time  `json:"expired_at" gorm:"column:expired_at"`
	Board     string     `json:"board" gorm:"column:board; type:int[]"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

type UserMatchCard struct {
	ID       int    `json:"uid" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
}
