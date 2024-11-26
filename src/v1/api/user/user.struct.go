package user

type UsersListResponse struct {
	ID       int    `json:"uid" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Win      int    `json:"win" gorm:"column:win"`
	Lose     int    `json:"lose" gorm:"column:lose"`
	Draw     int    `json:"draw" gorm:"column:draw"`
	Total    int    `json:"total" gorm:"column:total"`
}
