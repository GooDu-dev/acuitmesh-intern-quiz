package v1

import (
	"net/http"
	"time"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/api/test"
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
	testService []route
	userService []route
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
			Validation:  validator.NoValidation,
			Endpoint:    userEndpoint.GetUsersList,
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
	return ro
}
