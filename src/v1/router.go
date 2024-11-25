package v1

import (
	"net/http"
	"time"

	validator "github.com/GooDu-dev/acuitmesh-intern-quiz/src/v1/middlewares"
	"github.com/GooDu-dev/acuitmesh-intern-quiz/src/v1/services/test"
	"github.com/GooDu-dev/acuitmesh-intern-quiz/utils"

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
	for _, e := range r.searchService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	for _, e := range r.mapService {
		mainRoute.Handle(e.Method, e.Path, e.Validation, e.Endpoint)
	}
	return ro
}
