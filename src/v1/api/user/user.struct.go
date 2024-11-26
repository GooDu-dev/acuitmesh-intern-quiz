package user

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
	ID       int    `json:"uid" gorm:"column:id"`
	Username string `json:"username" gorm:"username"`
}
