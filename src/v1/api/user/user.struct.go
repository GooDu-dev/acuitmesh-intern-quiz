package user

import (
	"time"
)

type UsersListResponse struct {
	ID       int    `json:"uid" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Win      int    `json:"win" gorm:"column:win"`
	Lose     int    `json:"lose" gorm:"column:lose"`
	Draw     int    `json:"draw" gorm:"column:draw"`
	Total    int    `json:"total" gorm:"column:total"`
}

type UserHistoryResponse struct {
	Win   int           `json:"wins"`
	Lose  int           `json:"loses"`
	Draw  int           `json:"draw"`
	Total int           `json:"total"`
	Match []HistoryMath `json:"match"`
}

type UserStatistic struct {
	Win   int `json:"wins" gorm:"column:win"`
	Lose  int `json:"loses" gorm:"column:lose"`
	Draw  int `json:"draws" gorm:"column:draw"`
	Total int `json:"total" gorm:"column:total"`
}

type HistoryMath struct {
	ID     int      `json:"id"`
	Home   UserCard `json:"home"`
	Away   UserCard `json:"away"`
	Winner string   `json:"winner"`
}

type UserCard struct {
	ID        int       `json:"uid" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Token     string    `json:"token" gorm:"column:token"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"pwd"`
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
}

type UserCreateStruct struct {
	ID        int       `json:"id" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username"`
	Email     string    `json:"email" gorm:"column:mail"`
	Pwd       string    `json:"pwd" gorm:"column:pwd"`
	CreatedAt time.Time `json:"create_at" gorm:"column:created_at"`
}

type UserTokenStruct struct {
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	Token     string    `json:"token" gorm:"column:token"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
	OnUsed    string    `json:"on_used" gorm:"column:on_used"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type UserInviteStruct struct {
	HomeID      int        `json:"home_id" gorm:"column:home_id"`
	AwayID      int        `json:"away_id" gorm:"column:away_id"`
	StartTime   time.Time  `json:"start" gorm:"column:start_time_stamp"`
	ExpiredTime time.Time  `json:"expired" gorm:"column:expired_time_stamp"`
	IsAccept    string     `json:"is_accept" gorm:"column:is_accept"`
	Token       string     `json:"token" gorm:"column:token"`
	CreatedAt   *time.Time `json:"-" gorm:"column:created_at"`
}

type UserMatchStruct struct {
	HomeID int    `json:"home_id" gorm:"column:home_id"`
	AwayID int    `json:"away_id" gorm:"column:away_id"`
	Token  string `json:"token"  gorm:"column:token"`
	InviteID int `json:"-" gorm:"column:invite_id"`
}
