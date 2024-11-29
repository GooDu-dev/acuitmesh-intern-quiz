package v1

import (
	"net/http"
	"time"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/api/test"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/api/tictactoe"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/api/user"
	validator "github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/middlewares"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Path        string
	Validation  gin.HandlerFunc
	Endpoint    gin.HandlerFunc
}

type Router struct {
	testService      []route
	userService      []route
	tictactoeService []route
}

func (r Router) InitRouter() http.Handler {
	testEndPoint := test.NewEndPoint()
	r.testService = []route{
		{
			Name:        "[GET] : Ping",
			Description: "If works, 'Pong' will returns",
			Method:      http.MethodGet,
			Path:        "/test/ping",
			Validation:  validator.NoValidation,
			Endpoint:    testEndPoint.GetTestHealth,
		},
	}

	userEndpoint := user.NewEndpoint()
	r.userService = []route{
		{
			Name:        "[GET] Health Check",
			Description: "user service health check",
			Method:      http.MethodGet,
			Path:        "/user/health",
			Validation:  validator.NoValidation,
			Endpoint:    userEndpoint.GetUserHealth,
		},
		{
			Name:        "[GET] Get users list",
			Description: "get 50 users per list, get start index from params",
			Method:      http.MethodGet,
			Path:        "/users",
			Validation:  validator.AuthValidator,
			Endpoint:    userEndpoint.GetUsersList,
		},
		{
			Name:        "[GET] Get user statistic and history match",
			Description: "get user statistic and 20 each list of history match",
			Method:      http.MethodGet,
			Path:        "/user/dashboard",
			Validation:  validator.AuthValidator,
			Endpoint:    userEndpoint.GetUserDashboard,
		},
		{
			Name:        "[POST] User login with email and password",
			Description: "use email and password to login and return id and username back",
			Method:      http.MethodPost,
			Path:        "/auth/login",
			Validation:  validator.NoValidation,
			Endpoint:    userEndpoint.LoginUser,
		},
		{
			Name:        "[POST] User register with email, username, password",
			Description: "use email, username and password to register and return id and username back",
			Method:      http.MethodPost,
			Path:        "/auth/register",
			Validation:  validator.NoValidation,
			Endpoint:    userEndpoint.CreateUser,
		},
		{
			Name:        "[GET] Send invite to user",
			Description: "send invite to others player with his/her token",
			Method:      http.MethodGet,
			Path:        "/user/invite",
			Validation:  validator.AuthValidator,
			Endpoint:    userEndpoint.UserCreateInvite,
		},
		{
			Name:        "[GET] Accept the invitation",
			Description: "Accept the invitation via invitation token",
			Method:      http.MethodGet,
			Path:        "/user/invite/accept",
			Validation:  validator.AuthValidator,
			Endpoint:    userEndpoint.AcceptUserInvite,
		},
	}

	tictactoeEndpoint := tictactoe.NewEndpoint()
	r.tictactoeService = []route{
		{
			Name:        "[GET] Get match from token",
			Description: "get the user specified match info with token",
			Method:      http.MethodGet,
			Path:        "/tictactoe/match",
			Validation:  validator.AuthValidator,
			Endpoint:    tictactoeEndpoint.GetUserMatchInfo,
		},
		{
			Name:        "[PATCH] Set X,O to board",
			Description: "set player mark to board and refresh token",
			Method:      http.MethodPatch,
			Path:        "/tictactoe/match",
			Validation:  validator.AuthValidator,
			Endpoint:    tictactoeEndpoint.SetXOToBoard,
		},
	}

	ro := gin.Default()
	ro.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{utils.CONTENT_TYPE, utils.CONTENT_CODE, utils.ACCESS_CONTROL, utils.SOURCE_CONTROL},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	mainRoute := ro.Group(utils.PATH)
	for _, e := range r.testService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	for _, e := range r.userService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	for _, e := range r.tictactoeService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	return ro
}
