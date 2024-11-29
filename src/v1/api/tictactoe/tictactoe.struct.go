package tictactoe

import "github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/api/user"

type MatchInfo struct {
	HomePlayer user.UserCard `json:"home"`
	AwayPlayer user.UserCard `json:"away"`
	Board      BoardInfo     `json:"board"`
}

type BoardInfo struct {
	Data [9]int `json:"data"`
	Next string `json:"next"`
}
