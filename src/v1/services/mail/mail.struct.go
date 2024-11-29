package mail

import "fmt"

type MailUser struct {
	Username string `json:"username"`
	Email    string `json:"-"`
}

type MailContent struct {
	topic string
}

type MailMessage interface {
	ThHeader(string) string
	EnHeader(string) string
}

type UserMail interface {
	ThHeader(player1 string) string
	EnHeader(player1 string) string
	ThMessage(player1 string, player2 string, url string) string
	EnMessage(player1 string, player2 string, url string) string
}

type ThHeader func(string) string
type EnHeader func(string) string
type ThMessage func(string, string, string) string
type EnMessage func(string, string, string) string

func New(topic string) *MailContent {
	return &MailContent{
		topic: topic,
	}
}

func (mc *MailContent) ThHeader(player1 string) string {
	if Th, ok := mailThHeader[mc]; ok {
		return Th(player1)
	} else {
		return ""
	}
}
func (mc *MailContent) EnHeader(player1 string) string {
	if En, ok := mailEnHeader[mc]; ok {
		return En(player1)
	} else {
		return ""
	}
}
func (mc *MailContent) ThMessage(player1 string, player2 string, url string) string {
	if Th, ok := mailThMessage[mc]; ok {
		return Th(player1, player2, url)
	} else {
		return ""
	}
}
func (mc *MailContent) EnMessage(player1 string, player2 string, url string) string {
	if En, ok := mailEnMessage[mc]; ok {
		return En(player1, player2, url)
	} else {
		return ""
	}
}

var (
	UserInvite             UserMail = New("user_invite_another")
	UserWin                UserMail = New("user_win")
	UserLose               UserMail = New("user_lose")
	UserAcceptTheChallenge UserMail = New("user_accept_the_chanllenge")
)

var mailEnHeader = map[MailMessage]EnHeader{
	UserInvite:             func(player string) string { return fmt.Sprintf("User [%s] has invited you to battle!!", player) },
	UserWin:                func(player string) string { return fmt.Sprintf("You win, [%s]", player) },
	UserLose:               func(player string) string { return fmt.Sprintf("You lost, [%s]", player) },
	UserAcceptTheChallenge: func(player string) string { return fmt.Sprintf("User [%s] aceptted the challenge", player) },
}

var mailThHeader = map[MailMessage]ThHeader{
	UserInvite: func(player string) string {
		return fmt.Sprintf("ผู้ใช้ [%s] ได้เชิญคุณเข้าเล่นเกม", player)
	},
	UserWin:  func(player string) string { return fmt.Sprintf("คุณชนะ, [%s]", player) },
	UserLose: func(player string) string { return fmt.Sprintf("คุณแพ้, [%s]", player) },
	UserAcceptTheChallenge: func(player string) string {
		return fmt.Sprintf("ผู้ใช้ [%s] ยอมรับคำเชิญเข้าเล่นเกม", player)
	},
}

var mailThMessage = map[MailMessage]ThMessage{
	UserInvite: func(player1 string, player2 string, url string) string {
		return fmt.Sprintf("สวัสดีคุณ [%s] ผู้ใช้ [%s] ได้ทำการเชิญคุณเข้าสู่เกมผ่านลิ้ง [%s]", player1, player2, url)
	},
	UserWin: func(player1 string, player2 string, url string) string {
		return ""
	},
	UserLose: func(player1 string, player2 string, url string) string {
		return ""
	},
	UserAcceptTheChallenge: func(player1 string, player2 string, url string) string {
		return fmt.Sprintf("สวัสดีคุณ [%s] ผู้ใช้ [%s] ได้ตอบรับคำเชิญคุณเข้าสู่เกมผ่านลิ้ง [%s]", player1, player2, url)
	},
}

var mailEnMessage = map[MailMessage]EnMessage{
	UserInvite: func(player1 string, player2 string, url string) string {
		return fmt.Sprintf("็Hello [%s], user [%s] has invited you via [%s]", player1, player2, url)
	},
	UserWin: func(player1 string, player2 string, url string) string {
		return ""
	},
	UserLose: func(player1 string, player2 string, url string) string {
		return ""
	},
	UserAcceptTheChallenge: func(player1 string, player2 string, url string) string {
		return fmt.Sprintf("Hello [%s], user [%s] has accepted your challenge via [%s]", player1, player2, url)
	},
}
