package database

import "time"

type TicTacToeModel struct {
	InviteID          int             `json:"invite_id" gorm:"column:invite_id; index"`
	InvitationModelFK InvitationModel `gorm:"foreignKey:InviteID"`
	Token             string          `json:"token" gorm:"column:token; index"`
	IsEnd             string          `json:"is_end" gorm:"column:is_end; type:char(1)"`
	ExpiredAt         time.Time       `json:"expired_at" gorm:"column:expired_at; type:timestamptz"`
	Board             []int           `json:"board" gorm:"column:board; type:integer[]"`
	CreatedAt         time.Time       `json:"created_at" gorm:"column:created_at; type:timestamptz"`
	UpdatedAt         time.Time       `json:"updated_at" gorm:"column:updated_at; type:timestamptz"`
	DeletedAt         time.Time       `json:"deleted_at" gorm:"column:deleted_at; type:timestamptz"`
}

func (TicTacToeModel) TableName() string {
	return "tb_tictactoe"
}
